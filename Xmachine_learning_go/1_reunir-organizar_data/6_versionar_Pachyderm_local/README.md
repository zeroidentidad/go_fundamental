
# Instalación y deploy Pachyderm 1.9.x latest [23/01/2020]

## Doc: https://docs.pachyderm.com/latest/getting_started/local_installation/

## **Pasos en Ubuntu 18.04.3 LTS**

### **0** - Instalaciones previas de VirtualBox y Docker Engine en sus métodos de instalación de repositorio DEB, enlaces:

    - https://www.virtualbox.org/wiki/Linux_Downloads

    - https://docs.docker.com/install/linux/docker-ce/ubuntu/

* De acuerdo a preferencias sobre el paso **3**

### **1** - Para verificar si la virtualización es compatible con Linux, ejecutar el comando y verificar que la salida no esté vacía:

```shell
grep -E --color 'vmx|svm' /proc/cpuinfo
```

    Deberá mandar cadenas de valores dependiendo del numero de nucleos.

```shell
egrep -q 'vmx|svm' /proc/cpuinfo && echo yes || echo no
```
    Debera mandar 'yes' en caso de estar soportada    

### **2** - Instalación kubectl (Kubernetes CLI) con método de gestion automatica: https://kubernetes.io/docs/tasks/tools/install-kubectl/#install-kubectl-on-linux

- Se opto por el método de snap package manager:
```shell
snap install kubectl --classic
```
- Confirmar instalación:
```shell
kubectl version
```

    * La instalación usando la administración de paquetes nativos esta soportada hasta xenial, ref. comunidad: https://stackoverflow.com/questions/53068337/unable-to-add-kubernetes-bionic-main-ubuntu-18-04-to-apt-repository

 ### **3** - Instalación Minikube: https://kubernetes.io/docs/tasks/tools/install-minikube/

 - Se opto por el método de paquete para linux (https://minikube.sigs.k8s.io/docs/start/linux/): 

    Versión estable a elaboración de este README (*minikube_1.6.2.deb*)

 - Descarga e instalación:
```shell
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube_1.6.2.deb \
 && sudo dpkg -i minikube_1.6.2.deb
```
- Confirmar instalación, referencia driver a utilizar: https://minikube.sigs.k8s.io/docs/reference/drivers/
    
- = **Teniendo VirtualBox instalado** =
```shell
minikube start --vm-driver=virtualbox
```
- = **Hacer virtualbox el driver por default** =
```shell
minikube config set vm-driver virtualbox
```    
    
    Para solo usar: minikube start

    * Minikube también admite una opción --vm-driver=none que ejecuta los componentes de Kubernetes en el Host y no en una Maquina Virtual. Cuando se hace la instalación por paquetes apt de Docker, es cuando se puede utilizar el controlador none. En la instalación con Snap de docker no funcionan la opcion con minikube.

- = **Teniendo Docker instalado** =
```shell
sudo minikube start --vm-driver=none
```
- = **Hacer none sin driver por default** =
```shell
sudo minikube config set vm-driver none
```

    *- En caso de error de permisos usando --vm-driver=none optar por usar con virtualvox
    *- Hacer minikube delete y eliminar ~/.minikube (directorio de archivos de caché), creado incompleto

- Una vez finalice inicio de minikube, ejecutar el comando para verificar el estado del clúster:
```shell
minikube status
```

### **4** - Instalación pachctl (Pachyderm Client): https://docs.pachyderm.com/latest/getting_started/local_installation/#install-pachctl

- Descarga e instalación:
```shell
curl -o /tmp/pachctl.deb -L https://github.com/pachyderm/pachyderm/releases/download/v1.9.11/pachctl_1.9.11_amd64.deb && sudo dpkg -i /tmp/pachctl.deb
```

- Confirmar instalación:
```shell
pachctl version --client-only
```

- Deploy local:
```shell
pachctl deploy local
```

- Verificar estado de los pods de Pachyderm en deploy ejecutando varias veces kubectl get pods
```shell
kubectl get pods
NAME                     READY     STATUS              RESTARTS   AGE
dash-6c9dc97d9c-vb972    0/2       ContainerCreating    0          6m
etcd-7dbb489f44-9v5jj    1/1       Running              0          6m
pachd-6c878bbc4c-f2h2c   1/1       Running              0          6m
```
    * Cuando Pachyderm está listo para usarse, todos los pods de Pachyderm deben estar en el estado Running.
    Para ver detalladamente ejecutar:
```shell
kubectl get all
```    

- Ejecutar **pachctl version** para verificar que pachd hizo el deploy.
```shell
pachctl version
COMPONENT           VERSION
pachctl             1.9.11
pachd               1.9.11
```

- Utilizar el reenvío de puertos para acceder al panel de control Pachyderm.
```shell
pachctl port-forward
```
    *Este comando debe dejarse ejecutando continuamente y no sale a menos que lo interrumpa para terminar de usar la instancia del dashboard.

El resultado seria acceder como en la siguiente url: **http://localhost:30080**

- Alternativamente, configurar Pachyderm para conectarse directamente a la instancia de Minikube:

- -Obtener IP de Minikube:
```shell
minikube ip
```

- -Configurar Pachyderm para conectarse directamente a la instancia de Minikube:
```shell
pachctl config update context `pachctl config get active-context` --pachd-address=`num.ip.sin.comillas`:30650
```

El resultado seria acceder como en la siguiente url: **http://192.168.99.100:30080**

## Nota extra:
### Sobre la persistencia al iniciar instancia.

- Esto es principalmente una cuestión de cómo implementar kubernetes. Hasta donde sé, no hay una buena manera de persistir cuando se ejecuta minikube en modo vm, que es el valor predeterminado. No está destinado a cargas de trabajo de larga duración, sino más bien como un entorno de prueba/desarrollo. Se puede intentar ejecutar con --vm-driver=none (solo Linux). Usar asi funciona y persiste. También se podría considerar configurar kubernetes en un proveedor de la nube, o usar la oferta en la nube: hub.pachyderm.com aunque actualmente también está diseñado más como sandbox y tiene un tiempo de vencimiento establecido en clústeres.