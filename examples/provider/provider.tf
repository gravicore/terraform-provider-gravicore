terraform {
  required_providers {
    asm = {
      source  = "gravicore/gravicore"
      version = "1.0.1"
    }
  }
}

provider "gravicore" {
  region = "us-east-1"
}
