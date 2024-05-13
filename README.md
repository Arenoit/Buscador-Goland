# Buscador con Autocompletado
Indicaciones básicas, este miniproyecto esta hecho con Postgres de Docker, lenguaje de programación Golang
Esta versión diseñada esta hecha sin framworks con un manejo considerado de práctica
Autor: David Jiménez
## 👯 Contenido
Sistema que consta de una CRUD y buscador
- Motor de busqueda filtrado y Autocompletador.
- Ajax para busqueda en tiempo real
- Insertar,Editar y Eliminar simple sin asistencia de Javascript
- Base de Datos Postgres asistida por GORM (herramienta de asignación relacional de objetos)

[![Buscador.png](https://i.postimg.cc/ncxRHx3g/Buscador.png)](https://postimg.cc/8syBH2Qm)

## 💬 Código - Indicaciones
Se necesita de una terminal, instalar Goland previamente, necesario estar dentro de la carpeta
Comandos a ejecutar en Comand Promt o en PowerShell:

    cd C:/Directorio/ruta/carpeta
    go run ./app.go

Comandos para la instalacción de Posgrest en Docker:

    sudo docker run --name some-postgres -e POSTGRES_USER=arenoit POSTGRES_PASWORD=123 -p 5432:5432 -d postgres
    sudo docker exec -it some-postgres bash
    psql -U arenoit --password
    CREATE DATABASE buscador;
    \c buscador
    \d
    \d tasks
    SELECT * FROM tasks;

Hay que considerar que al poner el comando nos pedira la contraseña que creamos del contendor

    psql -U arenoit --password

En el caso que se salga y se quiera volver a entrar usar, las sentencias SQL necesitan (;)

    sudo docker start some-postgres
    sudo docker exec -it some-postgres bash
    psql -U arenoit --password
    \c buscador
    \d tasks
    SELECT * FROM tasks;

** Docker

Utilizar Docker para Windows, es necesario tener Hyper-V que solo soportan pocas computadoras para la Virtualización Anidada,
en el caso que la máquina que tenemos no lo soporte utilizar una máguina vritual normal como VirtualBox que necesita.
 - Adaptador Puente
 - etho

## 🧐 Detalles
La Virtualización Anidada nos permite utilizar Docker con nuestro propio pin 127.0.0.1 en vez de usar Máquinas Virtuales y realizar ajustes
Eso no significa que no se esta utilizando Linux, es Linux conectado directamente al PATH de Windows.

## ⚡ Tecnologías
Talk to me about
- Front-end **CSS, Javascript**
- Back-end **GO (Golang), Docker**
- Biblioteca o Herramienta para conexiones por ORM **GORM**
- Arquitectura de desarrollo **Sistema Base de Datos Relacionales**
- Implementación Back-end
- Aplicación Web **Diseño Responsivo**

---
⭐️ Nota
Solo es necesario crear la base de datos, las tablas se crean automáticamente por la aplicación web
