# Golang Backend-For-Frontend with gRPC

This project implements a cinema seating management system with **gRPC**, **GraphQL**, and **Kafka**. The system manages seat configurations, reservations, cancellations, and seat availability queries while enforcing manthattan distance rules.

## Overview

This system allows cinema operators to configure seating, manage seat reservations and cancellations, and query available seats while maintaining required social distancing. The system uses:

- **gRPC** for efficient inter-service communication.
- **GraphQL** to provide a flexible API for client queries.
- **Kafka** to publish events for reservations and cancellations, allowing for monitoring and event-driven architecture.

## Architecture

The system is built with three main components:

1. **gRPC Server**: Core service handling seat operations and applying distancing rules.
2. **GraphQL Server**: API layer allowing flexible client queries and mutations.
3. **Kafka**: Message broker for event streaming of reservation and cancellation actions.

[![](https://mermaid.ink/img/pako:eNqtVl1v2jAU_SuWpUqdRMsKG23zUKmiU_fQTitV94Q0GfsmserYmT-oUMV_nx0HCCOwoI0HlNicc-895zjkHVPFACfYwC8HksIdJ5kmxVQi_ymJtpzykkiLXgzo3dV7Tcr86QE9g5637WeT7-O9mxNg3Owu3xFLZsRA3InfJydorGTKM6e5zNCYSyhI3AqNnd3cbHeSoEdnieVKorM1EmocOtXqzfQQVcIV0l8UXP70rVjiBfgQWbfpPH9jkgSF6wZvpO3A2iA5a-m5bo-ueNmeXirhEvRSMmIBUUJzCO1so6vp9xCsJE7Qs1X6KLBHB8UTdEtfpXoTwDLYhiHjKAXjnV17NwEPmQfniOchtqtzEQcBEtRdiXuER4cZiLBVP-h2TrggMwGISIYeAaxBd9G70PXECaiT2sHGilHHwmylRuqEWGwoOnpaUdUzsAPoXUNbkCAMxPVvqjmy0ugHV8LXbQ7dedxYJFqWek4_86kJVZwkzRrzVQ22rlHbAJIdzlpD0VjHX5dKhmfE5vkQzpkQx4csAv8hYwcJ1hHjpsXILkmiFf1_SdK4pjo-SdvI7SQdPVZtVUtkpNqcnKPCQZuULel4cqAXIRub2AecORiSCuQVWGMakH2yj3Ogr2vVW6uFSIyrH3zldiNaRdB-wKzTMnIytDlTZkNZGRI5H7kxB_xtye892LYJOxj54E8yUml7TwdVqlptH-Vvftdy7IBxDxegC8KZf595DxxTbHMoYIoTf5n6NBs7xb16p7oJG7nPjUJvSgs2xVO59DzEWfW8kBQnVjvoYa1cluMkJV7mHnbVyarflP5Y_cK4PzbrRaEIA3_7ju2iDO9ZmVfMF4h_mGHdaeGXc2tLk_T7Yfs84zZ3s3Oqir7hLPevRvn8etQfDUZXZDCE0eWQfB4OGZ1dXF-lg08XKbv8eDEgeLlc_gZJqkL1?type=png)](https://mermaid.live/edit#pako:eNqtVl1v2jAU_SuWpUqdRMsKG23zUKmiU_fQTitV94Q0GfsmserYmT-oUMV_nx0HCCOwoI0HlNicc-895zjkHVPFACfYwC8HksIdJ5kmxVQi_ymJtpzykkiLXgzo3dV7Tcr86QE9g5637WeT7-O9mxNg3Owu3xFLZsRA3InfJydorGTKM6e5zNCYSyhI3AqNnd3cbHeSoEdnieVKorM1EmocOtXqzfQQVcIV0l8UXP70rVjiBfgQWbfpPH9jkgSF6wZvpO3A2iA5a-m5bo-ueNmeXirhEvRSMmIBUUJzCO1so6vp9xCsJE7Qs1X6KLBHB8UTdEtfpXoTwDLYhiHjKAXjnV17NwEPmQfniOchtqtzEQcBEtRdiXuER4cZiLBVP-h2TrggMwGISIYeAaxBd9G70PXECaiT2sHGilHHwmylRuqEWGwoOnpaUdUzsAPoXUNbkCAMxPVvqjmy0ugHV8LXbQ7dedxYJFqWek4_86kJVZwkzRrzVQ22rlHbAJIdzlpD0VjHX5dKhmfE5vkQzpkQx4csAv8hYwcJ1hHjpsXILkmiFf1_SdK4pjo-SdvI7SQdPVZtVUtkpNqcnKPCQZuULel4cqAXIRub2AecORiSCuQVWGMakH2yj3Ogr2vVW6uFSIyrH3zldiNaRdB-wKzTMnIytDlTZkNZGRI5H7kxB_xtye892LYJOxj54E8yUml7TwdVqlptH-Vvftdy7IBxDxegC8KZf595DxxTbHMoYIoTf5n6NBs7xb16p7oJG7nPjUJvSgs2xVO59DzEWfW8kBQnVjvoYa1cluMkJV7mHnbVyarflP5Y_cK4PzbrRaEIA3_7ju2iDO9ZmVfMF4h_mGHdaeGXc2tLk_T7Yfs84zZ3s3Oqir7hLPevRvn8etQfDUZXZDCE0eWQfB4OGZ1dXF-lg08XKbv8eDEgeLlc_gZJqkL1)

## Prerequisites

- **Go** (1.23 or newer)
- **Docker** (for Kafka setup)
- **Protocol Buffers** (for gRPC code generation)
- **Kafka Broker** (default setup with Docker Compose)

## Usage

- Go build

