import Ribbon from "./Ribbon.tsx"
import Network from "./Network.tsx"
import './App.css'
import { Image } from "./models/Image.tsx"
import { NetworkContainers } from "./models/Network.tsx"
import { JSX } from "react"

let rawdata = ``

function parseResponse(response: string): Map<string, NetworkContainers> {
  let images: Image[];
  try {
    images = JSON.parse(response) as Image[];
  } catch (e) {
    images = [];
  }

  let Networks: Map<string, NetworkContainers> = new Map<string, NetworkContainers>();
  
  for(let i = 0; i < images.length; i++) {
    if(Networks.has(images[i].Networks)) {
      Networks.get(images[i].Networks)?.Images.push(images[i]);
    } else {
      Networks.set(images[i].Networks, {networkName: images[i].Networks, Images: [images[i]]})
    }
  }

  return Networks;
}

function generateNetworkHTML(Networks: Map<string, NetworkContainers>): JSX.Element[] {

  let html: JSX.Element[] = [];

  Networks.forEach((value, key) => {
    console.log(`Key: ${key}, Value: ${value}`);
    html.push(<Network NetworkInfo={value} />)
  });

  return html
}

function App() {

  

  

  return (
    <>
      <Ribbon />
      {
        generateNetworkHTML(parseResponse(rawdata))
      }
    </>
  )
}

export default App
