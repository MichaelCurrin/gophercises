# Requerimiento Día 6

Dentro de esta carpeta "docs" se encuentra un archivo llamado `finalproject.postman_collection.json`, el mismo se puede importar desde postman. 

La finalidad de este es mostrar una serie de request con sus correspondientes query params y payloads, que la aplicación puede responder. Así también los distintos casos de uso que son solicitados en el ejercicio.


## Importante

Este requerimiento es libre y se realiza de forma individual, partiendo de la premisa de que el mismo debe de aportar un valor agregado a la aplicación, utilizando datos ya existentes o datos nuevos que se crean necesarios. Aclaración: no es necesario crear una aplicación nueva.


## User Story

|**[User Story Code: ml-parts-01]**|Horas estimadas|
| :- | :-: |
|**COMO** usuario de la plataforma **QUIERO** poder subscribirme a notificaciones de los pedidos por medio de email **PARA** poder verificar el estado de un pedido sin tener que consultar a la plataforma directamente **Nota:** Es necesario validar que los correos no sean de dominios descartables, esto se debe a que no queremos enviar emails innecesarios, tambien obtener datos del usuario (Ciudad, Pais, Moneda) que se registra por medio de su direccion ip para poder enviar emails personalizados.  **Nota Opcional:** Implementar cache para las request a sistemas externos. El envio de email tiene que ser asincronico . |**10**|
Pruebas de Usuario: usuario se subscribe correctamente, usuario obtiene informacion de ip, usuario obtiene informacion de moneda.|

## Representación JSON

Ejemplo de json de entrada y salida

`/api/v1/service/subscribe?email=[email]`

Devuelve datos del usuario subscripto

Parametros de entrada

|Nombre|Descripción|Obligatorio|
| :- | :- | :- |
|email|Direccion de email del usuario al cual se quiere subscribir a notificaciones|si|

Parametros de salida

|Nombre|Descripción|
| :- | :- |
|email|Email del usuario subscripto
|dateSubscribed|Fecha de subscripcion
|country|Pais del usuario
|ip|Direccion ip del usuario

`/api/v1/service/subscribe/me`

Parametros de salida

|Nombre|Descripción|
| :- | :- |
|email|Email del usuario subscripto
|dateSubscribed|Fecha de subscripcion
|country|Pais del usuario
|ip|Direccion ip del usuario

`/api/v1/service/email/{email}`

Devuelve informacion respecto al dominio del email.

Parametros de entrada

|Nombre|Descripción|Obligatorio|
| :- | :- | :- |
|email|Direccion de email a verificar el dominio|si|

Parametros de salida

|Nombre|Descripción|
| :- | :- |
|domain|Dominio del usuario subscripto
|disposable|Valor booleano que nos indica si el dominio es descartable

`/api/v1/service/ip/{ip}`

Devuelve informacion respecto a la ip solicitada.

Parametros de entrada

|Nombre|Descripción|Obligatorio|
| :- | :- | :- |
|ip|Direccion de ip que queremos obtener informacion|si|

Parametros de salida

|Nombre|Descripción|
| :- | :- |
|ip|ip del cual se obtuvo la informacion
|city|ciudad de la ubicacion de la ip
|region|region de la ip
|country|pais de la ip
|languages|lenguages del area de la ip
|currency|moneda del area de la ip

`/api/v1/service/currency/{currency}`

Devuelve informacion respecto a la moneda solicitada

Parametros de entrada

|Nombre|Descripción|Obligatorio|
| :- | :- | :- |
|currency|moneda de la cual queremos obtener informacion|si|

Parametros de salida

|Nombre|Descripción|
| :- | :- |
|currency|moneda de la cual se obtuvo la informacion
|rate|ciudad de la ubicacion de la ip

`/api/v1/email/{email}/{text}`

Envia un email al usuario

Parametros de entrada

|Nombre|Descripción|Obligatorio|
| :- | :- | :- |
|email|Usuario al cual se le quiere enviar un email|si|
|text|Mensaje|si

## Contratos referentes a la User Story

|HTTP|Plantilla URI|Descripción|US-code|
| :- | :- | :- | :- |
|POST|/api/v1/service/subscribe?email=[email]|Solo podra ser utilizado por un usuario con rol admin, mediante el cual podra subscribir a un usuario espesifico a que reciba notificacioens de sus pedidos |ml-parts-06|
|POST|`/api/v1/service/subscribe/me`|El usuario logeado actualmente se subscribira a notificaciones de sus pedidos |ml-parts-06|
|GET|`/api/v1/service/email/{email}`|Se obtendra informacion del dominio al cual pertenece el email. Solo podra ser utilizado por un usuario admin |ml-parts-06|
|GET|`/api/v1/service/ip/{ip}`|Se obtendra informacion de la ip solicitada. Solo podra ser utilizado por un usuario admin |ml-parts-06|
|GET|`/api/v1/service/currency/{currency}`|Se obtendra informacion de la moneda solicitada. Solo podra ser utilizado por un usuario admin |ml-parts-06|
|GET|`/api/v1/service/infoUser/{email}`|Se obtendra informacion de un usuario en particular . Solo podra ser utilizado por un usuario admin |ml-parts-06|
|GET|`/api/v1/service/infoUser/me`|Se obtendra informacion del usuario actualmente logeado . |ml-parts-06|
|POST|`/api/v1/service/email/{email}/{text`}|Se enviara un email al usuario indicado con el texto declarado. Solo podra ser utilizado por un usuario admin |ml-parts-06|


## Nota

Contemplar otros tipos de errores.

Trabajar con Access Token.
