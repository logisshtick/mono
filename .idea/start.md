# Start

## System Overview

```mermaid
graph LR

subgraph Client
  A[React] -->|HTTPS:4433| B1[API request]
  A -->|HTTPS:443| B2[Static request]
end

subgraph Server
  B1 -->|HTTPS| C1[nginx]
  B2 -->|HTTPS| C2[nginx]
  C1 -->|HTTP| C[API]
  C2 -->|HTTP| Static
end

subgraph Google APIs
  C -->|HTTPS| D[Google Maps Routes]
  C -->|HTTPS| E[Google Maps Places]
end
```

# End

# Tech Stack Start
|             | backend                 | frontend |
|-------------|-------------------------|----------|
| Language:   | golang + C/C++          | js/ts    |
| Frameworks: | net/http, encoding/json | react    |
| G-Wagon:    | yes                     | no       |


# Infrastructure

















