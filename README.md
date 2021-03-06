# grpc-rest-graphql-comparison

- Example gallery app

## FE

- React
  - Apollo
  - Axios

## BE

- Go (`:5000`)
  - GraphQL
  - REST
  - gRPC
- Node (`:5001`)
- Python (`:5002`)

## DB

- Redis

### Dev

- follow setup `/database/dev/redis-setup.md`
- 3 separate docker containers for redis master and replicas and 3 containers for sentinels

### Prod

- follow setup `/database/prod/redis-setup.md`
- kind-k8s cluster
  - 3 pods => master and two replicas
  - 3 sentinels => if master dies sentinel will promote one of the replicas to master

## Presentation

1. History => First there was REST...
2. Basics & Comparison of REST and GraphQL
3. gRPC how it works
4. Use cases
5. Pros & Cons
6. Comparison with other API paradigms (Webhook, Websocket, HTTP streaming, SOAP)

## API (Application Programming Interface)

- API is an interface trough which information and functionality can be shared between two services, where one service makes request and the other one provides response.

## REST (Representational State Transfer)

- returns JSON / XML
- basic CRUD operations on endpoints
- nonCRUD operations
  - Add flag in response
    - `PATCH /users/user-1`
    - `Body: {"archived": true}`
  - Add subresource as action's name
    - `PUT /users/user-1/archive`
  - Add query parameter for search
    - `PUT /users/user?id=1`

### Pros

- Standard method and arguments names
- Utilizes HTTP features
- Easy to maintain

### Cons

- Big payloads
- Multiple HTTP calls for resource and subresource

## GraphQL (Graph Query Language)

- returns JSON
- expose single endpoint
  - `app.com/graphql`
- developed in Facebook
- Only uses GET and POST requests
- Example apps => github, pinterest, airbnb

### Pros

- Define the exact data that is required
- Multiple nested data in single call, which results in quicker response time and smaller payload size
- Avoids versioning => has only one endpoint, in comparison with REST, where changing one endpoint can break whole app

### Cons

- Added complexity of handling the type of query that client constructs
- Optimizing performance => when working with multiple external clients their use cases may vary, so you have to find middleground for all of them
  - First client needs query for name and surname, while second client needs query for name, surname and age

## gRPC (grpc Remote Procedure Call)

- returns binary
- uses HTTP/2
- action functions on endpoints
- Originated in google
- Only uses GET and POST requests
- Example apps => slack

### Pros

- Easy to understand
- Lightweight payloads
- High performance

### Cons

- Not standardized so it can be difficult discover them
  - `app.com/api/user.getUsername`
- Limited standardization
- Can lead to function explosion as over time there will be more and more functions so it is important to keep track of all of them

## Conclusion

### REST

- Data driven
- Best for APIs that expose CRUD operations

### GraphQL

- Best for APIs that need querying flexibility

### gRPC

- Best for APIs exposing several actions instead of CRUD operations

### SOAP (Simple object access protocol)

- Function driven
- Before REST
- Messaging protocol specification for exchanging structured data
- main focus on easy communication between different platforms and languages

### Websocket

### Webhook

- Webhook is an event driven API.
- It is a service that allows one program to send data to another one as soon as a particular webhook event happens
- Client > Server communication with no response
- When user clicks on button event happens and data gets send to server instead of server continuously asking if button is clicked

### HTTP 1.1

- Using TCP connection
- With each new request comes new TCP connection

### HTTP 2.0

- Utilizes multiplexing to combine and compress requests into single TCP connection and giving them stream ID so after response is returned client can tell which request is which
- It pushes new responses into one TCP connetion if request requires more responses

GQL => jeden request kter?? vr??t?? v??ce dat
gRPC with HTTP 2.0 => v??ce request?? kombinovan??ch v multiplexovan??m TCP spojen??

### Slide Intro

Kdo z v??s u?? sly??el o gRPC (grpc remote procedure call) protokolu?
J?? jsem se o t??to technologii dozv??d??l kdy?? jsem se v??ce za??al zaj??mat o jazyk Go a sv??t krypta, mimochodem v??ele doporu??uju Go v??em backen????k??m, kdy?? u?? ne pro long-time use tak aspo?? na vyzkou??en?? a pochopen?? mo??nosti paraleln??ho programov??n??, jazyk se pou????v?? p??edev????m pro cloud engeneering a pro projekty kter?? si cht??j?? zachovat jednoduch?? memory management, rychlost a n??kter?? vlastnosti low level jazyk??. Zrovna kdy?? jsem pracoval na jednom zku??ebn??m projektu, kter?? vyu????val pr??v?? Go na Backendu a React na Frontendu tak m?? za??alo zaj??mat jak?? jsou nejlep???? zp??soby komunikace mezi r??zn??mi jazyky a platformami tak aby se dala jedna paradigma API reusnout na co nejv??ce servis??.

### Slide REST

P??i hled??n?? odpov??di jsem do??el k n??zoru ??e opravdu z??le???? na dan??m use casu, pokud n??m jde o jednoduch?? projekt se z??kladn??mi CRUD operacemi je tady REST. Jist?? v??ichni v??te jak funguje ale jen pro zopakov??n??:

- Definujeme v n??m samostatn?? endpointy pro ka??d?? resource
- Vyu????v?? standard?? HTTP metod (GET, POST, PUT, PATCH, ...)
- Nej??ast??ji pou????v?? JSON v n??kter??ch p????padech XML serializaci dat
- ??asto se dostaneme do situace kdy pracujeme s hromadou nepot??ebn??ch dat
- P??i async requestech a eventech nevyu????v?? plnou funkcionalitu HTTP2 k ??emu?? se dostanu pozd??ji

### Slide GraphQL

- Sta???? n??m jeden endpoint pro v??echny resources
- Vyu????v?? pouze http metod GET a POST
- Pou????v?? JSON serializaci dat
- Definujeme v n??m p??esnou strukturu dat kterou pot??ebujeme, co?? n??m dovoluje mnohon??sobn?? vno??en?? dat v jednom requestu narozd??l od RESTu kde pro ka??d?? resource d??l??me zvl?????? request
- Nejl??pe se pou????v?? na API, kter?? pot??ebuj?? flexibiln?? query. Nap????klad pokud bychom pou????vali n??jakou 3rd party API, kter?? pou????v?? GraphQL mohla by nastat situace kdy jeden klient pot??ebuje specifickou ????st dat a budeme muset pro ka??d??ho klienta vytv????et novou query

- Jsou tady je??t?? dal???? mo??nosti jako t??eba SOAP, Thrift, nebo JSON RPC, ale na ty se dnes nebudu zam????ovat

### Slide Evoluce HTTP

#### HTTP 0.9

- Prvn?? verze HTTP
- Podpora jedin?? metody GET
- Nepou????vala se pln?? funkcionalita URL
- P??es jedno TCP p??ipojen?? se mohl pos??lat pouze jeden soubour ve form??tu html `GET /index.html` vr??tilo `index.html`

#### HTTP 1.0

- Prvn?? stabiln?? verze
- P??ibyly funkce jako
  - status code
  - http headery
  - content type d??ky kter??mu se krom?? html dokument?? daly p??en????et i dokumenty jin??ho typu

#### HTTP 1.1

- Standardizovan?? verze
- Mo??nost transf??ru v??ce soubor?? v jednom TCP p??ipojen??
- encoding jazyk??

- mezit??m p??i??ly dal???? verze a funkcionality jako nap????klad autentika??n?? header nebo caching

#### HTTP 2.0

- Z??sadn?? zm??na pro TCP komunikaci kdy je d??ky multiplexingu mo??n?? v jednom TCP p??ipojen?? ud??lat v??ce request??, kompresnout je a p??idat jim stream ID pro pozd??j???? identifikac?? v response
- Funkcionalita server push, nap????klad kdy?? ud??l??me request na `index.html` a server si v??imne ??e na n??m z??vis?? `index.css` a `main.js` p??ipoj?? je do response, p??i ??patn?? konfiguraci nebo ??patn?? nastaven??m cachov??n?? v??ak m????e zp??sobit nadbyte??n?? data
- Binary serializace
- HTTPS by default
- Je pomalej???? pokud vytv??????me komunikaci se servisem kter?? pou????v?? HTTP 1.1

#### HTTP 3.0

- Zm??na na multiplexovan?? transfer dat pomoc?? QUIC p??es protokol UDP

### Slide gRPC

- Ka??d?? komunika??n?? protokol pot??ebuje client library, co?? je probl??m pro aplikace mimo prohl????e?? kde p??i ka??d?? nov?? verzi nebo featu??e mus??me manu??ln?? p??id??vat nov?? verze
<div style="text-align: center;">
  <img alt="TCP connection comparison" src="img/tcp-connection-comparison.png">
</div>

Uk??zat comparison HTTP 1.1 a HTTP 2.0 v developer tools network tabu
