package main

import "fmt"

type Producto struct {
	id       int
	nombre   string
	precio   float64
	cantidad int
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) *Producto {
	return &Producto{
		id:       id,
		nombre:   nombre,
		precio:   precio,
		cantidad: cantidad,
	}
}

//geters
func (p *Producto) GetID() int {
	return p.id
}

func (p *Producto) GetNombre() string {
	return p.nombre
}

func (p *Producto) GetPrecio() float64 {
	return p.precio
}

func (p *Producto) GetCantidad() int {
	return p.cantidad
}

//

//seters
func (p *Producto) SetPrecio(nuevoPrecio float64) {
	if nuevoPrecio >= 0 {
		p.precio = nuevoPrecio
	} else {
		fmt.Println("Error: El precio no puede ser negativo")
	}
}

func (p *Producto) SetCantidad(nuevaCantidad int) {
	if nuevaCantidad >= 0 {
		p.cantidad = nuevaCantidad
	} else {
		fmt.Println("Error: La cantidad no puede ser negativa")
	}
}

//

func (p *Producto) MostrarInfo() string {
	return fmt.Sprintf("Id: %d\n  Producto: %s\n  Precio: $%.2f\n  Cantidad: %d\n",
		p.id,
		p.nombre,
		p.precio,
		p.cantidad)
}

type Inventario struct {
	productos map[int]*Producto
}

func NuevoInventario() *Inventario {
	return &Inventario{
		productos: make(map[int]*Producto),
	}
}

func (i *Inventario) AgregarProducto(p *Producto) {
	i.productos[p.GetID()] = p
	fmt.Printf("Producto '%s' \n", p.GetNombre())
}

func (i *Inventario) EliminarProducto(id int) {
	if producto, existe := i.productos[id]; existe {
		delete(i.productos, id)
		fmt.Printf("Producto '%s' \n", producto.GetNombre())
	}
}

func (i *Inventario) BuscarProducto(id int) (*Producto, bool) {
	producto, existe := i.productos[id]
	fmt.Printf("Producto '%s' \n", producto.GetNombre())
	return producto, existe
}

func (i *Inventario) ListarProductos() {
	fmt.Println("**********Ineventario**********\n")
	for _, producto := range i.productos {
		fmt.Println(producto.MostrarInfo())
	}

}

func main() {
	inventario := NuevoInventario()

	producto1 := NuevoProducto(1, "Laptop", 450.99, 10)
	producto2 := NuevoProducto(2, "Telefono", 270.99, 20)
	producto3 := NuevoProducto(3, "Tablet", 320.99, 15)
	producto4 := NuevoProducto(4, "Pantalla", 150.99, 8)
	fmt.Println("**********Productos Agregados**********")
	inventario.AgregarProducto(producto1)
	inventario.AgregarProducto(producto2)
	inventario.AgregarProducto(producto3)
	inventario.AgregarProducto(producto4)
	fmt.Println("**********Productos Eliminados**********")
	inventario.EliminarProducto(1)
	fmt.Println("**********Productos Encontrados**********")
	inventario.BuscarProducto(2)
	inventario.ListarProductos()

}
