# Leal Backend

# Como inicializar el proyecto
1. Crea una base de datos en postgresql con el nombre leal.
2. Configura tu archivo .env definiendo el puerto en el que va escuchas el servidor y el dns de conexion a la base de datos.
3. Crea algun comercio en la base de datos.
4. Crea algunas sucursales en la base de datos.
5. Crea tu usuario en la base de datos.
6. Crea los siguientes beneficios descritos en el archivo tipo_beneficios.csv
7. Utiliza el endpoint para crear una campaña. Respeta las relaciones con el comercio y con la sucursal. Tambien no olvides el tipo de beneficio que definiste en el paso anterior.
8. Utiliza el endpoint para consultar las campañas de un comercio y una sucursal.
9. Utiliza el endpoint para crear una compra y con esto sumar puntos si aplica.
10. Si necesitas ver las sucursales, comercios y tipo de beneficios creados en la base de datos usa los endpoints relacionado a cada entidad.

# Getting Started
### Prerequisites
- Go 1.22 (debería seguir siendo compatible con versiones anteriores)

# Used Tools
1. <a HREF="https://github.com/gin-gonic/gin">Gin para manejar la capa de adaptadores entrantes. </a>
2. <a HREF="https://gorm.io/index.html"> Gorm para manejar la capa de adaptadores salientes. </a>
3. <a HREF="https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL"> PostgreSQL como base de datos. </a>
