package nvmeof

import (
    "fmt"
    "os/exec"
	"context"
	"github.com/netapp/trident/storage"
	"github.com/netapp/trident/storage_drivers"
	"bytes"
	"strconv"
	"strings"
)


const (
	StorageDriverName = "nvmeof"
)

func (d *NVMeoFDriver) Name() string {
	return StorageDriverName
}

type NVMeoFBackend struct {
    config NVMeoFConfig
}

func NewNVMeoFBackend(config NVMeoFConfig) (*NVMeoFBackend, error) {
	cmd := exec.Command("nvme", "discover", "-t", config.TransportType, "-a", config.TargetAddress, "-s", config.TargetPort)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return nil, fmt.Errorf("NVMe-oF discovery failed: %v, output: %s", err, output)
    }
    return &NVMeoFBackend{config: config}, nil
}

func (backend *NVMeoFBackend) Create(ctx context.Context, 
		volConfig *storage.VolumeConfig, storagePool storage.Pool, 
		volAttributes map[string]storage.AttributeRequest,
	) error{
    cmd := exec.Command("nvme", "create-ns", "--size", fmt.Sprintf("%d", volConfig.size))
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("error creating volume: %v, output: %s", err, output)
    }
    return nil
}

func (d *Driver) Initialize(
    ctx context.Context, driverContext storage.DriverContext, configJSON string,
    commonConfig *storage_drivers.CommonStorageDriverConfig, backendSecret map[string]string, backendUUID string,
) error {
    if err := json.Unmarshal([]byte(configJSON), &d.config); err != nil {
        return fmt.Errorf("unable to parse NVMe-oF configuration: %v", err)
    }

    if err := d.connectToNVMeTarget(); err != nil {
        return fmt.Errorf("unable to connect to NVMe-oF target: %v", err)
    }
    d.initializeVolumeTracking()

    d.backendUUID = backendUUID

    storage_drivers.Logc(ctx).WithFields(log.Fields{
        "driverName": commonConfig.StorageDriverName,
        "backendUUID": backendUUID,
    }).Info("Storage driver initialized")

    return nil
}

func (d *Driver) connectToNVMeTarget() error {
    cmd := exec.Command("nvme", "connect", "-t", d.config.TransportType, "-a", d.config.TargetAddress, "-s", d.config.TargetPort)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("error connecting to NVMe-oF target: %v, output: %s", err, string(output))
    }
    return nil
}

func (d *Driver) initializeVolumeTracking() {
	if d.volumes == nil {
        d.volumes = make(map[string]*Volume)
    }

    cmd := exec.Command("nvme", "list")
    output, err := cmd.CombinedOutput()
    if err != nil {
        log.Errorf("Failed to list NVMe namespaces: %v, output: %s", err, string(output))
        return
    }
    namespaces := parseNVMeListOutput(output)

    for _, ns := range namespaces {
        vol := &Volume{
            Name:      ns.Name,
            Size:      ns.Size,
            Namespace: ns.NamespaceID,
        }
        d.volumes[ns.Name] = vol
    }

    log.Infof("Successfully initialized volume tracking with %d volumes.", len(d.volumes))

}

func parseNVMeListOutput(output []byte) []Namespace {
	var namespaces []Namespace
	lines := strings.Split(string(output), "\n")

	for _, line := range lines {
		fields := strings.Fields(line)

		
		if len(fields) >= 3 {
			size, err := strconv.ParseInt(fields[2], 10, 64)
			if err != nil {
\				continue
			}
			namespace := Namespace{
				NamespaceID: fields[0],
				Name:        fields[1],
				Size:        size,
			}
			namespaces = append(namespaces, namespace)
		}
	}

	return namespaces
}

func (backend *NVMeoFBackend) Destroy(
    ctx context.Context, 
    volConfig *storage.VolumeConfig) error {
    cmd := exec.Command("nvme", "delete-ns", volConfig.name)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("error deleting volume: %v, output: %s", err, output)
    }
    return nil
}

func (backend *NVMeoFBackend) Resize(ctx context.Context, volConfig *VolumeConfig, sizeBytes uint64) error {
    cmd := exec.Command("nvme", "resize", volConfig.name, fmt.Sprintf("--size=%d", volConfig.size))
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("error resizing volume: %v, output: %s", err, output)
    }
    return nil
}

func (backend *NVMeoFBackend) AttachVolume(volumeName string, hostName string) error {
    cmd := exec.Command("nvme", "connect", "--volume", volumeName, "--host", hostName)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("error attaching volume to host: %v, output: %s", err, output)
    }
    return nil
}

func (backend *NVMeoFBackend) DetachVolume(volumeName string) error {
    cmd := exec.Command("nvme", "disconnect", "--volume", volumeName)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("error detaching volume: %v, output: %s", err, output)
    }
    return nil
}

func init() {
    storage_drivers.RegisterDriver(nvmeof.StorageDriverName, NewNVMeoFBackend)
}
