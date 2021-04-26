pipeline {
    environment {
        registry = 'nha/shiba_auth'
        // registryCredential = 'dockerHubCredentials'
        dockerImage = ''
    }
    agent any
    triggers {
        pollSCM '* * * * *'
    }

    // tools {
    //     maven 'M3'
    // }
    stages {
        stage('Build Application') {
            steps {
                echo '=== Building Shiba Application ==='
                sh 'docker build -t shiba-auth-app .'
            }
        }
        // stage('Test Application') {
        //     steps {
        //         echo '=== Testing Petclinic Application ==='
        //         sh 'mvn test'
        //     }
        // //       post {
        // //        always {
        // //        junit 'target/surefire-reports/*.xml'
        // //    }
        // //           }
        // }
        stage('Build Docker Image') {
            steps {
                echo '=== Building Petclinic Docker Image ==='
                script {
                    dockerImage = docker.build registry + ":${env.BUILD_NUMBER}"
                }
            }
        }
        stage('Push Docker Image') {
            //        when {
            //            branch 'master'
            //            }
            steps {
                echo '=== Pushing Petclinic Docker Image ==='
                script {
                    // docker.withRegistry('', registryCredential ) {
                    //     dockerImage.push()
                    // }
                    docker.withRegistry("https://720480159010.dkr.ecr.ap-southeast-1.amazonaws.com/shiba_auth", "ecr:ap-southeast-1:aws_credential") {
                        docker.image($registry).push()
                    }
                }
            }
        }
        stage('Remove local images') {
            steps {
                echo '=== Delete the local docker images ==='
                sh("docker rmi -f $registry:$BUILD_NUMBER || :")
            }
        }
    }
}
