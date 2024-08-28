terraform {
  required_providers {
    gravicore = {
      source  = "gravicore/gravicore"
      version = "1.0.1"
    }
  }
}

provider "gravicore" {
  region = "us-east-1"
}
