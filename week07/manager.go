package week06

import (
	"context"
)

// Manager is responsible for receiving product orders
// and assigning them to employees. Manager is also responsible
// for receiving completed products, and listening for errors,
// from employees. Manager takes products that have been built
// by employees and returns them to the customer as a CompletedProduct.
type Manager struct {
	// non-exported fields (PRIVATE)
	// !YOU MAY NOT ACCESS THESE FIELDS IN YOUR TESTS!
	completed chan CompletedProduct
	errs      chan error
	jobs      chan *Product
	quit      chan struct{}
	stopped   bool
}

// NewManager will create a new Manager.
// This function should ALWAYS be used to
// create a new Manager.
func NewManager() *Manager {
	return &Manager{
		completed: make(chan CompletedProduct),
		errs:      make(chan error),
		jobs:      make(chan *Product),
		quit:      make(chan struct{}),
	}
}

// Start will create new employees for the given count,
// and start listening for jobs and errors.
// Managers should be stopped using the Stop method
// when they are no longer needed.
func (m *Manager) Start(ctx context.Context, count int) error {

	if count <= 0 {
		return ErrInvalidEmployeeCount(count)
	}

	go func() {
		<-ctx.Done()
		m.Stop()
	}()

	for i := 0; i < count; i++ {

		e := Employee(i + 1)
		go e.work(ctx, m)
	}

	return nil
}

// Assign will assign the given products to employees
// as employeess become available. An invalid product
// will return an error.
func (m *Manager) Assign(products ...*Product) error {
	if m.stopped {
		return ErrManagerStopped{}
	}

	// loop through each product and assign it to an employee
	for _, p := range products {
		// validate product
		if err := p.IsValid(); err != nil {
			return err
		}

		// assign product to employee
		// this will block until an employee becomes available
		if !m.stopped {
			m.Jobs() <- p
		}

	}

	return nil
}

// Complete will wrap the employee and the product into
// a CompletedProduct. The will be passed down the Completed()
// channel as soon as a listener is available to receive it.
// Complete will error if the employee is invalid or
// if the product is not built.
func (m *Manager) Complete(e Employee, p *Product) error {
	// validate employee
	if err := e.IsValid(); err != nil {
		return err
	}

	// validate product is built
	if err := p.IsBuilt(); err != nil {
		return err
	}

	cp := CompletedProduct{
		Employee: e,
		Product:  *p, // deference pointer to value type ype t
	}

	// Send completed product to Completed() channel
	// for a listener to receive it.
	// This will block until a listener is available.
	if !m.stopped {
		m.completedCh() <- cp
	}

	return nil
}

// completedCh returns the channel for CompletedProducts
func (m *Manager) completedCh() chan CompletedProduct {
	return m.completed
}

// Completed will return a channel that can be listened to
// for CompletedProducts.
// This is a read-only channel.
func (m *Manager) Completed() <-chan CompletedProduct {
	return m.completedCh()
}

// Done will return a channel that will be closed
// when the manager has stopped.
// Employees should listen to this channel to know
// when to stop listening for jobs.
func (m *Manager) Done() <-chan struct{} {
	return m.quit
}

// Jobs will return a channel that can be listened to
// for new products to be built.
func (m *Manager) Jobs() chan *Product {
	return m.jobs
}

// Errors will return a channel that can be listened to
// and can be used to receive errors from employees.
func (m *Manager) Errors() chan error {
	return m.errs
}

// Stop will stop the manager and clean up all resources.
func (m *Manager) Stop() {
	if m.stopped {
		return
	}

	m.stopped = true

	// close all channels
	close(m.quit)
	close(m.jobs)
	close(m.errs)
	close(m.completed)
}

func Run(ctx context.Context, emp int, count int, products ...*Product) ([]CompletedProduct, error) {

	// TODO: implement this function
	// This function should run the manager with the given products
	// and return the results.

	//start a new manager
	m := NewManager()
	defer m.Stop()

	//start with the number of employees
	err := m.Start(ctx, emp)
	if err != nil {
		return nil, err
		// m.Errors() <- err
	}

	//assign products to distribute jobs
	if !m.stopped {
		go func() {

			err = m.Assign(products...)
			if err != nil {
				m.Errors() <- err
			}
		}()
	}
	//storage for all completed product
	var cp []CompletedProduct

	go func() {
		for p := range m.Completed() {
			cp = append(cp, p)

			if len(cp) == count {
				m.Stop()
			}
		}

	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case e := <-m.Errors():
		m.Stop()
		return nil, e
	case <-m.Done():

	}

	return cp, nil
}
