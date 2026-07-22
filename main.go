package main

import "fmt"

// FUNCIONES

// 1. Función que nos da el precio de un producto
func obtenerPrecio(producto string) float64 {
	if producto == "laptop" {
		return 800.0
	} else if producto == "celular" {
		return 300.0
	} else if producto == "audifonos" {
		return 50.0
	}
	return 0.0 // Si no existe el producto
}

// 2. Función que calcula el impuesto IVA
func calcularIVA(subtotal float64) float64 {
	return subtotal * 0.15
}

// 3. Función que aplica un descuento si la compra es mayor a $500
func calcularDescuento(subtotal float64) float64 {
	if subtotal >= 500.0 {
		return subtotal * 0.10 // 10% de descuento
	}
	return 0.0
}

// PROGRAMA PRINCIPAL

func main() {
	var cliente string
	var subtotal float64 = 0.0
	var opcion int = 1

	fmt.Println("=== SISTEMA DE E-COMMERCE ===")
	fmt.Print("Ingrese el nombre del cliente: ")
	fmt.Scanln(&cliente)

	// Bucle para pedir productos repetidamente
	for opcion == 1 {
		var producto string
		var cantidad int

		fmt.Println("\nProductos disponibles: laptop ($800), celular ($300), audifonos ($50)")
		fmt.Print("Escriba el producto que desea comprar: ")
		fmt.Scanln(&producto)

		precio := obtenerPrecio(producto)

		if precio > 0 {
			fmt.Print("Ingrese la cantidad: ")
			fmt.Scanln(&cantidad)

			montoItem := precio * float64(cantidad)
			subtotal = subtotal + montoItem
			fmt.Printf("-> Agregado! Subtotal actual: $%.2f\n", subtotal)
		} else {
			fmt.Println("-> Producto no encontrado, intente de nuevo.")
		}

		fmt.Print("\n¿Desea agregar otro producto? (1 = Sí, 0 = No): ")
		fmt.Scanln(&opcion)
	}

	// Si compró algo, calculamos el total final
	if subtotal > 0 {
		descuento := calcularDescuento(subtotal)
		subtotalConDescuento := subtotal - descuento
		iva := calcularIVA(subtotalConDescuento)
		totalFinal := subtotalConDescuento + iva

		fmt.Println("\n=================================")
		fmt.Println("       RESUMEN DE FACTURA        ")
		fmt.Println("=================================")
		fmt.Printf("Cliente:          %s\n", cliente)
		fmt.Printf("Subtotal Bruto:   $%8.2f\n", subtotal)
		fmt.Printf("Descuento:       -$%8.2f\n", descuento)
		fmt.Printf("IVA (15%%):        +$%8.2f\n", iva)
		fmt.Println("---------------------------------")
		fmt.Printf("TOTAL A PAGAR:    $%8.2f\n", totalFinal)
		fmt.Println("=================================")
	} else {
		fmt.Println("\nNo se realizó ninguna compra.")
	}
}
