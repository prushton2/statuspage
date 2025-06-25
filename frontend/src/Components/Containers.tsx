import { JSX } from "react";
import "./Containers.css"
import { NetworkContainers } from "../models/Network";

export const Containers = ({Networks}: {Networks: NetworkContainers[] | undefined}) => {
    if(Networks == undefined) {
        return <>Loading...</>
    }
    let messages: {
        "health: starting": [string, string],
        "healthy":          [string, string],
        "unhealthy":        [string, string]
    } = {
        "health: starting": ["Service is Starting", "blue"],
        "healthy":          ["Service is Healthy", "green"],
        "unhealthy":        ["Service is Unhealthy", "red"]
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
            <div className="network-image-list">
                {Network.Images.map((container, index) => (
                    <div key={index} className="docker-image">
                        <p className="docker-image-status-dot">
                            {getServiceDot(container.Health)}
                        </p>
                        <h3 className="docker-image-title">
                            {container.Name}
                        </h3>

                        <p className="docker-image-element">
                            {container.Status}
                        </p>
                        
                        <p className="docker-image-element">
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