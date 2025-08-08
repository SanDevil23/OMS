# Order Management Service (OMS)

> Architecture


                            +-----------------------+
                            |     API Gateway       |   <-- Unified entrypoint for all services
                            |  (e.g., Kong, NGINX,  |
                            |   Go custom gateway)  |
                            +----------+------------+
                                        |
            +--------------------------+----------------------------+
            |                          |                            |
            +--v--+                  +----v----+                  +----v----+
            | MS1 |                  |  MS2     |                 |  MS3     |  <-- Go Microservices
            +-----+                  +---------+                 +----------+
            |                          |                            |
            +--v-------------------------v----------------------------v--+
            |            Service Discovery / API Registry (Optional)     |
            +------------------------------------------------------------+

            +------------------------------------------------------------+
            |        API Portal / UI (e.g., Swagger UI / Redoc)          |
            |    - Fetches OpenAPI specs from microservices              |
            |    - Displays docs and allows API testing on 1 page        |
            +------------------------------------------------------------+

