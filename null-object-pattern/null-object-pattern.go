package null_object_pattern

type AbstractCustomer interface {
	IsNull() bool
	GetName() string
}

type RealCustomer struct {
	Name string
}

func (r *RealCustomer) IsNull() bool {
	return false
}

func (r *RealCustomer) GetName() string {
	return r.Name
}

type NullCustomer struct {
}

func (n *NullCustomer) IsNull() bool {
	return true
}

func (n *NullCustomer) GetName() string {
	return "null customer"
}

type CustomerFactory struct {
	customers map[string]AbstractCustomer
}

func (c *CustomerFactory) Init() {
	c.customers = make(map[string]AbstractCustomer)
}

func (c *CustomerFactory) AddCustomer(customer AbstractCustomer) {
	c.customers[customer.GetName()] = customer
}

func (c *CustomerFactory) GetCustomer(name string) AbstractCustomer {
	if nil == c.customers[name] {
		return &NullCustomer{}
	}
	return c.customers[name]
}
