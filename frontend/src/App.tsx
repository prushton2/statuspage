import './App.css'
import { useEffect, useState } from "react";
import Containers from './Components/Containers.tsx';
import Ribbon from "./Components/Ribbon.tsx"
import { NetworkContainers } from './models/Network.tsx';
import { Image } from './models/Image.tsx';
import { getContainerInfo } from './API.tsx';
import { HealthSummary } from './models/HealthSummary.tsx';

function App() {
  const [networkContainers, setNetworkContainers] = useState<NetworkContainers[]>()
  const [healthSummary, setHealthSummary] = useState<HealthSummary>({healthy: 0, unhealthy: 0, starting: 0})

  useEffect(() => {
    async function init() {
      let containerInfo = await getContainerInfo()
      let networkContainers = convertToNetworkContainers(containerInfo)
      setHealthSummary(getHealthSummary(containerInfo))
      setNetworkContainers(networkContainers)
    }
    init()
  }, [])


  return (
    <>
      <Ribbon healthSummary={healthSummary}/>
      <Containers Networks={networkContainers}/>
    </>
  )
}

export default App


function convertToNetworkContainers(images: Image[]): NetworkContainers[] {
  let networks: Map<string, NetworkContainers> = new Map<string, NetworkContainers>

  for(let i = 0; i < images.length; i++) {
    if(images[i].Name[0] == ".") {
      continue
    }

    let item = networks.get(images[i].Network)

    if(item == undefined) {
      networks.set(images[i].Network, {networkName: images[i].Network, Images: []} as NetworkContainers)
      item = networks.get(images[i].Network)
    }

    item!.Images.push(images[i])
    networks.set(images[i].Network, item!)
  }

  let middleSegment: NetworkContainers[] = []
  let endSegment: NetworkContainers[] = []

  let endElements = ["syncthing_default", "statuspage_default", "host"]

  for(let i = 0; i < endElements.length; i++) {
    if (networks.get(endElements[i]) != undefined) {
      endSegment.push(networks.get(endElements[i])!)
      networks.delete(endElements[i])
    }
  }

  networks.forEach((value: NetworkContainers) => {
    middleSegment.push(value)
  })

  middleSegment.push(...endSegment)

  return middleSegment
}

function getHealthSummary(images: Image[]): HealthSummary {
  let summary: HealthSummary = {healthy: 0, unhealthy: 0, starting: 0}

  images.forEach((image) => {
    switch(image.Health) {
      case "healthy":
        summary.healthy++
        break;
      case "unhealthy":
        summary.unhealthy++;
        break;
      case "health: starting":
        summary.starting++;
        break;
    }
  })

  return summary
}