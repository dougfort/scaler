package methods

type serviceData struct {
	instanceCount uint32
	// TODO: resources
}

type platformData struct {
	services map[string]serviceData
}

func newPlatform(serviceIDs []string) platformData {
	var p platformData

	p.services = make(map[string]serviceData)
	for _, d := range serviceIDs {
		p.services[d] = serviceData{}
	}

	return p
}

func (p *platformData) enumerateServices() []string {
	var serviceIDs []string

	for serviceID := range p.services {
		serviceIDs = append(serviceIDs, serviceID)
	}

	return serviceIDs
}

// getServiceState returns false if the service does not exist on
// this platform. That is different from having zero instances.
func (p *platformData) getServiceState(
	serviceID string,
) (serviceData, bool, error) {
	data, ok := p.services[serviceID]
	if !ok {
		return data, false, nil
	}
	return data, true, nil
}

// setServiceState returns false if the deployment does not exist on
// this platform. That is different from having zero instances.
func (p *platformData) setServiceState(
	serviceID string,
	data serviceData,
) (bool, error) {
	_, ok := p.services[serviceID]
	if !ok {
		return false, nil
	}

	p.services[serviceID] = data

	return true, nil
}
