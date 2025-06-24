import './App.css'
import { useEffect, useState } from "react";
import Containers from './Components/Containers.tsx';
import Ribbon from "./Components/Ribbon.tsx"
import { NetworkContainers } from './models/Network.tsx';
import { Image } from './models/Image.tsx';
import { getContainerInfo } from './API.tsx';

function App() {
  const [networkContainers, setNetworkContainers] = useState<NetworkContainers[]>()

  useEffect(() => {
    async function init() {
      let images: Image[] = await getContainerInfo()
      let networks: Map<string, NetworkContainers> = new Map<string, NetworkContainers>

      for(let i = 0; i < images.length; i++) {
        let item = networks.get(images[i].Network);

        if(item == undefined) {
          networks.set(images[i].Network, {networkName: images[i].Network, Images: []} as NetworkContainers)
          item = networks.get(images[i].Network)
        }

        item!.Images.push(images[i])
        networks.set(images[i].Network, item!);
      }

      let networksList: NetworkContainers[] = []
      networks.forEach((value: NetworkContainers, key: string) => {
        networksList.push(value);
      })

      setNetworkContainers(networksList)
    }
    init()
  }, [])


  return (
    <>
      <Ribbon />
      <Containers Networks={networkContainers}/>
    </>
  )
}

export default App
