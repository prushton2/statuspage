import { JSX } from "react";
import "./Containers.css"
import { NetworkContainers } from "../models/Network";

export const Containers = ({Networks}: {Networks: NetworkContainers[] | undefined}) => {
    if(Networks == undefined) {
        return <>Loading...</>
    }
    let messages: {
        starting:  [string, string],
        healthy:   [string, string],
        unhealthy: [string, string]
    } = {
        starting:  ["Service is Starting", "blue"],
        healthy:   ["Service is Healthy", "green"],
        unhealthy: ["Service is Unhealthy", "red"]
    };
    
    const getServiceDot = (health: string): JSX.Element => {

        if(health == "") {
            return <></>
        }
        
        let healthInfo: [string, string] = messages[health as keyof typeof messages];
        
        return <span className={`dot ${healthInfo[1]}`} title={healthInfo[0]}/>
    }

    function getNetwork(Network: NetworkContainers) {
        return <div className="network-container">
            <h2 className="network-title">{Network.networkName}</h2>
            <div className="network-list">
                {Network.Images.map((container, index) => (
                    <div key={index} className="network-item">
                        {getServiceDot(container.Health)}
                        <h3 className="network-item-element">
                            {container.Name}
                        </h3>

                        <p className="network-item-element">
                            {container.Status}
                        </p>
                        
                        <p className="network-item-element">
                            {container.Size}
                        </p>
                    </div>
                ))}
            </div>
        </div>
    }

    return <>
        {Networks.map((v: NetworkContainers) => getNetwork(v))}
    </>
}

export default Containers;