# **Requerimiento Día 6**
**Especificaciones del Requerimiento**

**Requerimientos US: ml-parts-06**

[**Consultar pedidos de concesionario**](#_heading=h.2s8eyo1)

[**Descripción de parámetros de entrada y respuestas**](#_heading=h.gjdgxs)[](#_heading=h.gjdgxs)

[**User Story**](#_heading=h.3dy6vkm)[](#_heading=h.3dy6vkm)

[**Representación JSON**](#_heading=h.1t3h5sf)[](#_heading=h.1t3h5sf)

[**Contratos referentes a la User Story**](#_heading=h.4d34og8)[](#_heading=h.4d34og8)

**Importante:** 

Este requerimiento es libre y se realiza de forma individual, partiendo de la premisa de que el mismo debe de aportar un valor agregado a la aplicación, utilizando datos ya existentes o datos nuevos que se crean necesarios. Aclaración: no es necesario crear una aplicación nueva.


















**User Story**


 **[User Story Code: ml-parts-01]**|Horas estimadas|
| :- | :-: |
|**COMO** usuario de la plataforma **QUIERO** poder subscribirme a notificaciones de los pedidos por medio de email **PARA** poder verificar el estado de un pedido sin tener que consultar a la plataforma directamente **Nota:** Es necesario validar que los correos no sean de dominios descartables, esto se debe a que no queremos enviar emails innecesarios, tambien obtener datos del usuario (Ciudad, Pais, Moneda) que se registra por medio de su direccion ip para poder enviar emails personalizados.  **Nota Opcional:** Implementar cache para las request a sistemas externos. El envio de email tiene que ser asincronico . |**10**|
Pruebas de Usuario|

**Representación JSON**


Ejemplo de json de entrada y salida


**

**Contratos referentes a la User Story**


|HTTP|Plantilla URI|Descripción|US-code|
| :- | :- | :- | :- |
|POST|/api/v1/service/subscribe?email=[email]|Solo podra ser utilizado por un usuario con rol admin, mediante el cual podra subscribir a un usuario espesifico a que reciba notificacioens de sus pedidos |ml-parts-06|
|POST|/api/v1/service/subscribe/me|El usuario logeado actualmente se subscribira a notificaciones de sus pedidos |ml-parts-06|
|GET|/api/v1/service/email/{email}|Se obtendra informacion del dominio al cual pertenece el email. Solo podra ser utilizado por un usuario admin |ml-parts-06|
|GET|/api/v1/service/ip/{ip}|Se obtendra informacion de la ip solicitada. Solo podra ser utilizado por un usuario admin |ml-parts-06|
|GET|/api/v1/service/currency/{currency}|Se obtendra informacion de la moneda solicitada. Solo podra ser utilizado por un usuario admin |ml-parts-06|
|GET|/api/v1/service/infoUser/{email}|Se obtendra informacion de un usuario en particular . Solo podra ser utilizado por un usuario admin |ml-parts-06|
|GET|/api/v1/service/infoUser/me|Se obtendra informacion del usuario actualmente logeado . |ml-parts-06|
|POST|/api/v1/service/email/{email}/{text}|Se enviara un email al usuario indicado con el texto declarado. Solo podra ser utilizado por un usuario admin |ml-parts-06|


**Nota:** 

Contemplar otros tipos de errores. 

Trabajar con Access Token 


