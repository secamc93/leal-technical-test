# ---------------------------------------------------------
# Etapa de construcci├│n:
# Compila el binario est├ítico usando Go.
# ---------------------------------------------------------

# Usamos una imagen de Go para compilar el binario.
FROM golang:1.22-alpine3.20 AS builder

# Establecer un directorio de trabajo para la construcci├│n.
WORKDIR /app

# Copiar los archivos de m├│dulos
COPY go.mod go.sum ./

# Descargar los m├│dulos necesarios.
RUN go mod download

# Copiar el resto del c├│digo fuente del proyecto
COPY . .

# Compilar el binario est├ítico.
RUN go build -o leal-technical-test .

# ---------------------------------------------------------
# Etapa final: 
# Preparar la imagen de ejecuci├│n con el binario compilado.
# ---------------------------------------------------------
# Usamos una imagen de Alpine para reducir el tama├▒o de la imagen final.
FROM alpine:3.20
LABEL maintainer="Sebastian camacho <secamc_93@hotmail.com>" \
	  version="1.0" \
	  description="Imagen Docker para leal-technical-test"

# Establecer un directorio de trabajo para la ejecuci├│n.
WORKDIR /app

# Copiar el binario compilado de la etapa de construcci├│n.
COPY --from=builder /app/leal-technical-test /app/leal-technical-test

# Copiar los archivos de configuraci├│n y certificados.
COPY .env /app/.env
# COPY ./deployment/certs /app/deployment/certs

# Exponer el puerto en el que la aplicaci├│n escuchar├í.
EXPOSE 60000

# Ejecutar el binario.
CMD ["/app/SimonBK_Listener"]
