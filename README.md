# Database

Se utiliza como fuente de datos un archivo excel suministrado por los profesores de Digital House, el mismo para poder ser consumido por la aplicación se lo somete a un proceso ETL, es decir en una primera instancia se lo descarga a local para poder seguir con el siguiente proceso. En la etapa de transformación se parsea la información del usuario a información útil para nuestro dominio ejemplos de esto es un &quot;SI&quot; a un true, y se crean objetos en memoria de estos conceptos. Para la etapa de Load lo que se realiza almacenar estos objetos en un archivo de formato json el cual finalmente será consumido por la aplicación.

Para poder realizar todo este proceso, se debe ejecutar el siguiente archivo:

![alt text](https://i.postimg.cc/RhZwNK43/Screen-Shot-2021-03-08-at-11-31-45-AM.png)

Cabe mencionar que el ejecutar este archivo va a reemplazar nuestra base de datos de productos haciendo que todas las operaciones realizadas sobre este serán borradas volviendo al archivo original. Se recomienda probar a lo último debido a que cambiaran todas las id de productos y las request template del postman varias ya no serán válidas.

# Request Template

Dentro de esta carpeta &quot;docs&quot; se encuentra un archivo llamado &quot;desafiospring.postman\_collection.json&quot;, el mismo se puede importar desde postman. La finalidad de este es mostrar una serie de request con sus correspondientes query params y payloads, que la aplicación puede responder. Así también los distintos casos de uso que son solicitados en el ejercicio.

Una visualización de la carpeta que tendrían que ver en la sección collections del postman:

![alt text](https://i.postimg.cc/jCVqGYgs/Screen-Shot-2021-03-08-at-11-37-27-AM.png)

# Aclaracion

En esta sección se expondrán distintos agregados que no eran parte de la consigna.

## Transacciones

Dentro de este concepto se guardará registro de todas las operaciones de compras realizadas, llamase operación de compra a aquella en la que el usuario quiere hacer definitiva una compra directamente mediante la url ```api/v1/purchase-request```; o al finalizar un carrito de compra ```api/v1/shopping-cart/purchase-request/finisher```. Sus atributos son :

- Un id para poder identificar de manera unívoca una transacción en particular.
- Un state que puede tomar los valores de &quot;successful&quot; o &quot;fail&quot;, estos valores como bien indica su nombre es cuando falla o sale exitosa una transacción.
- El usuario que realizó la compra será identificado como guest user en caso de que no se haya logueado o caso contrario tendrá la información del usuario.
- Logs, es en esta propiedad donde se mostrará más en detalle qué es lo que sucedió dentro de la transacción, tanto para un caso de éxito como para uno de falla.
- Un ticket.

Ahora pasamos a definir cómo puede fallar una transacción. Al momento de realizar una compra por cada producto de la misma se van a ir ejecutando la operación de compra. Con las siguientes validaciones:

- La cantidad solicitada debe ser menor o igual a la que se posee actualmente en stock.
- Las cantidades solicitadas no pueden ser cero.

Si una transacción dentro de la operación de compra falla. Entonces se realizará un rollback de todas las ejecutadas hasta el momento y no se producirá efecto sobre nuestra base de datos (&quot;Database.json&quot;).

Caso contrario que todas las transacciones suceden de forma exitosa, producirá el efecto de reducción de stock en nuestra base de datos.

También estas transacciones no solo se registran en memoria, si no que se almacenan en su propio archivo/base de datos , &quot;Transactions.json&quot;

Para poder visualizar las transacciones se puede hacer una petición get a la siguiente url:

&quot;api/v1/transactions&quot;, soporta distintos filtros como id, state y today. Aclaración: state, responde a si la transacción es &quot;fail&quot; o &quot;successful&quot; y today con un valor true o false, para solicitar solo las transacciones de hoy o todas las que no son de hoy.

![alt text](https://i.postimg.cc/0QFJ23W4/Screen-Shot-2021-03-08-at-8-20-39-PM.png)

![alt text](https://i.postimg.cc/NjKqRzpD/Screen-Shot-2021-03-08-at-8-10-42-PM.png)

Ejemplo de log por falla de stock

![alt text](https://i.postimg.cc/J44WznX9/Screen-Shot-2021-03-09-at-10-05-40-AM.png)

Ejemplo de transacción exitosa con su correspondiente commit

![alt text](https://i.postimg.cc/GtjrSTyj/Screen-Shot-2021-03-09-at-10-06-58-AM.png)

## Filtros/Sorters

La aplicación permite aplicar &quot;n&quot; filtros y &quot;m&quot; sorters. La forma de agregarlos es creando la clase filtro u order y almacenando en un json la siguiente información:
```
[{

"id": x, "clazz" : y

}]
```
Donde la variable &quot;x&quot; es el nombre del queryParam que va a filtrar u ordenar.

Donde la variable &quot;y&quot; es el nombre de la clase correspondiente a este filtro u orden.

Actualmente los conceptos que tienen filtro y orden son: Articles.

Conceptos con filtros: Articles, Transactions, Clients

## Persistencia

Se persisten los datos de distintas entidades, como ya mencionamos transacciones y productos( Transactions.json y Database.json respectivamente). También se persisten &quot;Clients.json&quot; y por último pero no menos importante sesiones, las cuales se expondrán en la siguiente sección, &quot;Sessions.json&quot;.

## Sesiones

Las sesiones son utilizadas para poder obtener un carrito de compras previo, la misma está implementada en memoria con su correspondiente persistencia en &quot;Sessiones.json&quot;

El flujo de utilización de una sesión es mediante una request a

```api/v1/shopping-cart/purchese-request```, la primera vez que se envíe esta request vendrá sin cookie, una vez que el backend verifique esto se le asignará una sesión y dentro de la cookie se guardará esta.

![alt text](https://i.postimg.cc/Df1zX880/Screen-Shot-2021-03-08-at-12-06-49-PM.png)

De esta manera la próxima vez que hagamos una request solicitando agregar otros productos al carrito, ésta se enviará con la cookie de tal forma que el backend pueda identificar una sesión y actualice el carrito agregando estos nuevos productos.

Al hacer una request a la url ```api/v1/shopping-cart/purchase-request/finisher```, en este momento nosotros queremos finalizar la compra del carrito por esto deberíamos tener una cookie previa indicando nuestra sesión, si no se envia esto la request fallará. Una vez que la compra se realice, el carrito se vacía para que la persona pueda seguir comprando.

Lo mismo sucede con la url ```api/v1/shopping-cart/purchase-request/clean``` que se utiliza para limpiar un carrito de compras de una sesión en particular para esto es necesario este dato en la cookie.

## Registro de usuarios

Se agregó al registro de usuarios en el payload el campo password, este mismo será almacenado dentro de ```Clients.json``` junto con los demás datos de usuario, pero nunca se expondrá hacia afuera el password será de uso interno, esta encriptado usando BCrypt.

![alt text](https://i.postimg.cc/Bvt6LpKf/Screen-Shot-2021-03-08-at-8-14-37-PM.png)

## Login

Un usuario mediante un post a la ruta ```api/v1/login``` y un payload indicando dni y password, podrá loguearse dentro de la plataforma y tendrá su sesión disponible. ![alt text](https://i.postimg.cc/8zZcz9sb/Screen-Shot-2021-03-08-at-8-16-15-PM.png)

## Swagger

Se utilizó para poder documentar la api, y para poder visualizar esta documentación se utilizó la dependencia swagger-ui la cual se puede visualizar mediante
```api/v1/swagger```

## Logs de excepciones

Se implementó un logeo de todas nuestras excepciones de dominio lo que nos brinda información para la toma de decisiones. Estos vendrán ordenados de los más recientes a los más antiguos.

![alt text](https://i.postimg.cc/dVsDvwQ1/Screen-Shot-2021-03-09-at-3-39-44-PM.png)
