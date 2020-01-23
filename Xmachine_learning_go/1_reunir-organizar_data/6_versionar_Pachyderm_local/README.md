
# Instalacion y deploy Pachyderm 1.9.x latest

## Doc: https://docs.pachyderm.com/latest/getting_started/local_installation/

## **Pasos en Ubuntu 18.04.3 LTS**

**0** - Instalaciones previas de VirtualBox y Docker Engine en sus metodos de instalacion de repositorio DEB, enlaces:

    - https://www.virtualbox.org/wiki/Linux_Downloads

    - https://docs.docker.com/install/linux/docker-ce/ubuntu/

**1** - Para verificar si virtualización es compatible con Linux, ejecutar el comando y verificar que la salida no esté vacía:

    ```shell
    grep -E --color 'vmx|svm' /proc/cpuinfo
    ```

    Deberá mandar cadenas de valores dependiendo del numero de nucleos.