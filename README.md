# calculadora-avanzada

## Descripción
El programa debe permitir realizar operaciones matemáticas básicas y avanzadas. Entre las que pueden ser:
* aritmetica basicas (sumas, restas, multiplicaciones, divisiones, raices cuadradas, potencias, etc)
* aritmetica avanzada (logaritmos, senos, cosenos, tangentes, etc).
* algebra (ecuaciones de primer y segundo grado, matrices, etc)

Será construido en lenguaje Go para el BackEnd, no contará con una interfaz gráfica, sino que será una API REST que permita realizar las operaciones matemáticas usando el protocolo HTTP y Postman.

El programa será desplegado en un servidor de Heroku que tiene como rama principal la rama main.

## Requisitos
* El programa debe permitir realizar operaciones matemáticas básicas y avanzadas.
* El programa debe permitir realizar operaciones de algebra.
* El programa debe permitir realizar operaciones de geometría.
* El programa debe contar con un sistema de logeo usando contraseña y usuario.
* *Cada función o metodo debe contar con su prueba unitaria.*

## Go
Utilizaremos de base el framework [Gin](https://gin-gonic.com/docs/quickstart/) y [GORM](https://gorm.io/docs/index.html) para la conexión a la base de datos.

## Base de datos
EL programa debe contar con una base de datos que permita almacenar datos de los usuarios y que podamos logearnos con un usuario y contraseña. Será construida usando:
* PostgreSQL
* DiagramDB https://dbdiagram.io/home 

## Como arrancar el proyecto
1. Clonar el repositorio
2. Instalar las dependencias con el comando `go mod tidy`
3. Crear cada operación dentro de la paquetería correspondiente (aritmetica, algebra, geometria, ...) con la estructura de paquetes que se muestra a continuación.
- handler (contiene la interacción con el cliente)
- service (contiene la lógica de negocio)
- repository (en caso de ser necesario)
- spec (contiene la estructura de datos)
4. El programa se lanzará con el comando `go run cmd/main.go`

## Como contribuir
1. Crear una rama con el nombre de la funcionalidad que se va a implementar.
2. Crear un pull request con la funcionalidad implementada.
3. Esperar a que el pull request sea aceptado.

### Proceso para crear una rama
1. git checkout -b nombre_de_la_rama (crea la rama y te cambia a ella)
2. git add . (añade los cambios)
3. git commit -m "mensaje" (crea el commit)
4. git push origin nombre_de_la_rama (sube los cambios a la rama)

### Proceso para crear un pull request
1. Ir a la rama en la que se ha trabajado.
2. Hacer click en el botón "Compare & pull request".
3. Seleccionar la rama a la que se quiere hacer el pull request (preproduction).
4. Seleccionar al revisor del pull request.
5. Hacer click en el botón "Create pull request".
6. Esperar a que el pull request sea aceptado.