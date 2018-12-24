pipeline {
  agent any

  stages {
    stage("checkout") {
      steps {
        dir("tmp") {
          git changelog: false, poll: true, url: 'https://github.com/skeletonlander/docker-helloservice.git', branch: 'master'
        }
      }
    }
    stage("build") {
      steps {
        dir("tmp") {
          sh 'docker build'
        }
      }
    }
  }

  post {
    always {
      archive "target/**/*"
    }
  }
}
