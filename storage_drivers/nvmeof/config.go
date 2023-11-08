package nvmeof

const NVMeoFVersion = "1.0"

type NVMeoFConfig struct {
    Version           string `json:"version"`
    TargetAddress     string `json:"targetAddress"`
    TargetPort        string `json:"targetPort"`
    NamespaceID       string `json:"namespaceID"`
    TransportType     string `json:"transportType"`
    DiscoveryProtocol string `json:"discoveryProtocol"`
}