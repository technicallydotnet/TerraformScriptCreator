package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {

	var vmNumber, webAppNumber, acrNumber, aksNumber int
	var osparam, location, mainTF, variablesTF, providersTF, spec, vmSpec, webAppSpec, acrSpec, aksSpec string
	//var vmSpecs, webAppSpecs, acrSpecs, aksSpecs []string
	flag.IntVar(&vmNumber, "vmNumber", 0, "Number of virtual machines to create")
	flag.IntVar(&webAppNumber, "webAppNumber", 0, "Number of web apps to create")
	flag.IntVar(&acrNumber, "acrNumber", 0, "Number of ACR resources to create")
	flag.IntVar(&aksNumber, "aksNumber", 0, "Number of AKS resources to create")
	flag.StringVar(&osparam, "os", "windows", "Operating System of virtual machines to use")
	flag.StringVar(&location, "location", "West Europe", "Operating System of virtual machines to use")
	flag.StringVar(&spec, "spec", "low", "Level to use. eg- low, medium, high, or another spec specified")

	// Code to get spec from files (with custom specs set in file)
	/*vmSpecFile, vmerr := os.Open("vm-spec.txt") // the file is inside the local directory
	webAppSpecFile, waerr := os.Open("webapp-spec.txt")
	acrSpecFile, acrerr := os.Open("acr-spec.txt")
	aksSpecFile, akserr := os.Open("aks-spec.txt")
	if vmerr != nil || waerr != nil || acrerr != nil || akserr != nil {
		fmt.Println("Error")
	}


	defer vmSpecFile.Close()
	defer webAppSpecFile.Close()
	defer acrSpecFile.Close()
	defer aksSpecFile.Close()

	vmScanner := bufio.NewScanner(vmSpecFile)
	webAppScanner := bufio.NewScanner(webAppSpecFile)
	acrScanner := bufio.NewScanner(acrSpecFile)
	aksScanner := bufio.NewScanner(aksSpecFile)

	for vmScanner.Scan() {
		vmSpecs = append(vmSpecs, vmScanner.Text())
	}
	for webAppScanner.Scan() {
		webAppSpecs = append(webAppSpecs, vmScanner.Text())
	}
	for acrScanner.Scan() {
		acrSpecs = append(acrSpecs, vmScanner.Text())
	}
	for aksScanner.Scan() {
		aksSpecs = append(aksSpecs, aksScanner.Text())
	}*/

	if spec == "low" {
		vmSpec = "Standard_B1s"
		webAppSpec = "S1"
		acrSpec = "Basic"
		aksSpec = "B4ms"
	} else if spec == "medium" {
		vmSpec = "Standard_B2s"
		webAppSpec = "S1"
		acrSpec = "Standard"
		aksSpec = "DS4s_v4"
	} else {
		vmSpec = "Standard_D3s_v3"
		webAppSpec = "P2V2"
		acrSpec = "Premium"
		aksSpec = "E4ds_v4"
	}

	providersTF = `terraform {
    required_providers {
      azurerm = {
        source = "hashicorp/azurerm"
        version = "2.90.0"
      }
    }
  }
  
  provider "azurerm" {
    # Configuration options
  }`
	mainTF = "resource \"azurerm_resource_group\" \"rg\" {\nname     = \"example-resources\"  location = \"" + location + "\"}\n"
	var maintemp, variablestemp string
	for i := 0; i < vmNumber; i++ {
		if osparam == "windows" {
			maintemp = "resource \"azurerm_network_interface\" \"windows_nic_" + fmt.Sprint(i) + "\" {name                = var.windows_nic_" + fmt.Sprint(i) + "  location            = azurerm_resource_group.rg.location\n  resource_group_name = azurerm_resource_group.rg.name\n  ip_configuration {\n    name                          = \"internal\"\n    subnet_id                     = azurerm_subnet.example.id\n    private_ip_address_allocation = \"Dynamic\"\n  }}resource \"azurerm_windows_virtual_machine\" \"example\" {\n  name                = var.windows-machine-" + fmt.Sprint(i) + "\n  resource_group_name = azurerm_resource_group.rg.name\n  location            = azurerm_resource_group.rg.location\n  size                = \"" + vmSpec + "\"\n  admin_username      = \"adminuser\"\n  admin_password      = \"P@$$w0rd1234!\"\n  network_interface_ids = [\n    azurerm_network_interface.windows_nic_" + fmt.Sprint(i) + ".id,\n  ]\n  os_disk {\n    caching              = \"ReadWrite\"\n    storage_account_type = \"Standard_LRS\"\n  }\n  source_image_reference {\n    publisher = \"MicrosoftWindowsServer\"\n    offer     = \"WindowsServer\"\n    sku       = \"2016-Datacenter\"\n    version   = \"latest\"\n  }\n}\n"
			variablestemp = "variable \"windows_nic_" + fmt.Sprint(i) + "\" {\n    type = string\n}\nvariable \"windows_nic_" + fmt.Sprint(i) + "\" {\n    type = string\n}\n"
		} else {
			maintemp = "resource \"azurerm_network_interface\" \"linux_nic" + fmt.Sprint(i) + "\" {\n  name                = var.linux_nic_" + fmt.Sprint(i) + "\n  location            = azurerm_resource_group.rg.location\n  resource_group_name = azurerm_resource_group.rg.name\n  ip_configuration {\n    name                          = \"internal\"\n    subnet_id                     = azurerm_subnet.example.id\n    private_ip_address_allocation = \"Dynamic\"\n  }\n}\nresource \"azurerm_linux_virtual_machine\" \"linux_machine_" + fmt.Sprint(i) + "\" {\n  name                = var.linux_machine_" + fmt.Sprint(i) + "\n  resource_group_name = azurerm_resource_group.rg.name\n  location            = azurerm_resource_group.rg.location\n  size                = \"" + vmSpec + "\"\n  admin_username      = \"adminuser\"\n  network_interface_ids = [\n    azurerm_network_interface.linux_nic_" + fmt.Sprint(i) + ".id,\n  ]  admin_ssh_key {\n    username   = \"adminuser\"\n    public_key = file(\"~/.ssh/id_rsa.pub\")\n  }\n  os_disk {\n    caching              = \"ReadWrite\"\n    storage_account_type = \"Standard_LRS\"\n  }\n  source_image_reference {\n    publisher = \"Canonical\"\n    offer     = \"UbuntuServer\"\n    sku       = \"16.04-LTS\"\n    version   = \"latest\"\n  }\n}\n"
			variablestemp = "variable \"linux_nic_" + fmt.Sprint(i) + "\" {\n    type = string\n}\nvariable \"linux_nic_" + fmt.Sprint(i) + "\" {\n    type = string\n}\n"
		}
		mainTF += maintemp
		variablesTF += variablestemp
	}
	for i := 0; i < webAppNumber; i++ {
		maintemp := "resource \"azurerm_app_service_plan\" \"app_service_plan_" + fmt.Sprint(i) + "\" {\n  name                = var.app_service_plan_" + fmt.Sprint(i) + "\n  location            = azurerm_resource_group.rg.location\n  resource_group_name = azurerm_resource_group.rg.name\n  sku {\n    tier = \"Standard\"\n    size = \"" + webAppSpec + "\"\n  }\n}\nresource \"azurerm_app_service\" \"app_service_" + fmt.Sprint(i) + "\" {\n  name                = var.app_service_" + fmt.Sprint(i) + "\n  location            = azurerm_resource_group.rg.location\n  resource_group_name = azurerm_resource_group.rg.name\n  app_service_plan_id = azurerm_app_service_plan.example.id\n}\n"
		variablestemp := "variable \"app_service_plan_" + fmt.Sprint(i) + "\" {\n    type = string\n}\nvariable \"app_service_" + fmt.Sprint(i) + "\" {\n    type = string\n}\n"
		mainTF += maintemp
		variablesTF += variablestemp
	}
	for i := 0; i < vmNumber; i++ {
		maintemp := "resource \"azurerm_container_registry\" \"acr_" + fmt.Sprint(i) + "\" {\n  name                = \"acr_" + fmt.Sprint(i) + "\"\n  resource_group_name = azurerm_resource_group.rg.name\n  location            = azurerm_resource_group.rg.location\n  sku                 = \"" + acrSpec + "\"\n  admin_enabled       = false\n  georeplications {\n    location                = \"East US\"    zone_redundancy_enabled = true    tags                    = {}\n  }\n  georeplications {\n    location                = \"westeurope\"    zone_redundancy_enabled = true    tags                    = {}\n  }\n}\n"
		variablestemp := "variable \"acr_" + fmt.Sprint(i) + "\" {\n    type = string\n}\n"
		mainTF += maintemp
		variablesTF += variablestemp
	}
	for i := 0; i < vmNumber; i++ {
		maintemp := "resource \"azurerm_kubernetes_cluster\" \"aks_" + fmt.Sprint(i) + "\" {\n    name                = \"aks_" + fmt.Sprint(i) + "\"\n    location            = azurerm_resource_group.rg.location\n    resource_group_name = azurerm_resource_group.rg.name\n    dns_prefix          = \"aks" + fmt.Sprint(i) + "\"\n      default_node_pool {\n      name       = \"default\"\n      node_count = 1\n      vm_size    = \"" + aksSpec + "\"\n    }\n      identity {\n      type = \"SystemAssigned\"\n    }\n}\n"
		variablestemp := "variable \"aks_" + fmt.Sprint(i) + "\" {\n type = string\n}\n"
		mainTF += maintemp
		variablesTF += variablestemp
	}

	// Write all files at the end of the program
	ioutil.WriteFile("main.tf", []byte(mainTF), 0644)
	ioutil.WriteFile("variables.tf", []byte(variablesTF), 0644)
	ioutil.WriteFile("providers.tf", []byte(providersTF), 0644)
	fmt.Println("Finished creating terraform")
}
