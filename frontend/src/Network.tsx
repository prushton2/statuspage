import { JSX } from "react";
import "./Network.css"
import { NetworkContainers } from "./models/Network";

export const Network = ({NetworkInfo}: {NetworkInfo: NetworkContainers}) => {
    let messages: {
        starting:  [string, string],
        healthy:   [string, string],
        unhealthy: [string, string]
    } = {
        starting:  ["Service is Starting", "blue"],
        healthy:   ["Service is Healthy", "green"],
        unhealthy: ["Service is Unhealthy", "red"]
    };


    const formatPorts = (ports: string) => {
        return ports.split(", ").map((port, index) => (
            <span key={index}>
                {port}
                <br />
            </span>
        ));
    };
    
    const getServiceDot = (status: string): JSX.Element => {
        let health;

        try {
            health = status.split("(")[1].slice(0,-1);
        } catch (e) {
            return <span className="dot" />
        }

        let healthInfo: [string, string] = messages[health as keyof typeof messages];
        
        return <span className={`dot ${healthInfo[1]}`} title={healthInfo[0]}/>
    }

    return <>
        <div className="network-container">
            <h2 className="network-title">{NetworkInfo.networkName}</h2>
            <div className="network-list">
                {NetworkInfo.Images.map((container, index) => (
                    <div key={index} className="network-item">
                        {getServiceDot(container.Status)}
                        <h3 className="network-item-element">
                            {container.Names}
                        </h3>

                        <p className="network-item-element">
                            {container.Status}
                        </p>
                        
                        <p className="network-item-element">
                            {NetworkInfo.networkName != "host" ? formatPorts(container.Ports) : <></>}
                        </p>
                    </div>
                ))}
            </div>
        </div>
    </>
}

export default Network;