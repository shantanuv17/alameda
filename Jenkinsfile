node('go11') {
  stage('checkout') {
        git branch: 'master', url: "https://gitlab.prophetservice.com/ProphetStor/alameda.git"
  }
  stage("Build Operator") {
    sh """
      export GOROOT=/usr/local/go
      export GOPATH=/go/src/workspace
      mkdir -p /go/src/workspace/src/prophetstor.com
      mv ${env.WORKSPACE} /go/src/workspace/src/prophetstor.com/alameda
      cd /go/src/workspace/src/prophetstor.com/alameda/operator
      make manager
    """
  }
  stage("Build Datahub") {
    sh """
      export GOROOT=/usr/local/go
      export GOPATH=/go/src/workspace
      cd /go/src/workspace/src/prophetstor.com/alameda/datahub
      pwd
      make datahub
    """
  }
  stage("Test Operator") {
    sh """
      export GOROOT=/usr/local/go
      export GOPATH=/go/src/workspace
      cd /go/src/workspace/src/prophetstor.com/alameda/operator
      make test
    """
  }
  stage("Test Datahub") {
    sh """
      export GOROOT=/usr/local/go
      export GOPATH=/go/src/workspace
      cd /go/src/workspace/src/prophetstor.com/alameda/datahub
      curl -s https://codecov.io/bash | bash -s - -t ee1341b8-56f3-4319-8146-afc464130075
    """
  }
}
