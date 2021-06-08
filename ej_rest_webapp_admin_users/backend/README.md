# Backend: *API admin users*

- Bajar dependencias: `go mod vendor`

- Configuraciones necesarias: `.env`

- Contenedor DB: `docker-compose up -d`

- **Importar permisos DB**: `./permissions.sh`

- Para iniciar con hot-reload: `air` [Binario: https://github.com/cosmtrek/air]

- Ejecución preparando entorno: `./air.sh`

- Build web frontend: `../frontend/build/`