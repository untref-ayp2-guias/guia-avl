# Guía 14
## Implementación de las diapositivas

En la siguiente carpeta se encuentra la implementación del código suministrado en la clase:

- `/avltree` código del árbol AVL, el nodo AVL y los iteradores
- `/stack`, código de una pila sobre slices (se usa para implementar el iterador InOrder)
- `/queue`, código de una cola sobre slices (se usa para implementar un iterador por niveles)

## Ejercicios

### Ejercicio 1
En lápiz y papel dibujar cada paso. Crear un árbol AVL vacío e insertar: 10, 50, 30, 40, 9, 8, 11, 12, 13, 14, 15. Indicando a cada paso si se realizan rotaciones. Comparar los resultados obtenidos modificando el main

### Ejercicio 2
Escribir un iterador PreOrder.
> Pista: Al crear el iterador apilar la raiz
> Al devolver el siguiente desapilar y devolver el elemento del tope. Si ese nodo tiene hijo izquierdo apilarlo y si tiene hijo derecho apilarlo

### Ejercicio 3
Escribir un iterador PosOrder.
