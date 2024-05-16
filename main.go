package main

import (
	"context"
	"fmt"
	"log"
	"kafka-connect/msk" 
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func main() {
	region := "eu-west-1"
	clusterArn := "arn:" // ARN real de tu clúster de MSK

	// Crea la configuración de AWS incluyendo las credenciales explícitamente
	awsConfig, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			"",            // Access Key ID
			"", // Secret Access Key
			"", // Session Token, si es necesario; de lo contrario, dejar como cadena vacía
		)),
	)
	if err != nil {
		log.Fatalf("Error al cargar la configuración de AWS: %v", err)
	}

	// Crea una instancia de MskCluster con la configuración de AWS
	mskCluster := msk.NewMskClusterWithClientConfig(clusterArn, msk.SaslIam, awsConfig)

	// Obtén la configuración para conectarte a MSK
	opts, err := mskCluster.Config()
	if err != nil {
		log.Fatalf("Error al obtener la configuración de MSK: %v", err)
	}

	// Mostrar las opciones para verificar que todo está correcto (opcional)
	fmt.Println("Opciones de configuración de Kafka:", opts)
}
