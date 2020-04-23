package domain

type DatastoreAG struct {
	ID string
	Name string
	Path string
	Hostname string
	Port int
}

func NewDatastoreAG() *DatastoreAG {
	return &DatastoreAG{

	}
}

func (d *DatastoreAG) GetID() string {
	return d.ID
}

func (d *DatastoreAG) GetName() string {
	return d.Name
}

func (d *DatastoreAG) GetPath() string {
	return d.Path
}

func (d *DatastoreAG) GetHostname() string {
	return d.Hostname
}

func (d *DatastoreAG) GetPort() int {
	return d.Port
}
