import "./Network.css"
import { NetworkContainers } from "./models/Network";

export const Network = ({NetworkInfo}: {NetworkInfo: NetworkContainers}) => {

    const formatPorts = (ports: string) => {
        return ports.split(", ").map((port, index) => (
            <span key={index}>
                {port}
                <br />
            </span>
        ));
    };

    return <>
        <div className="network-container">
            <h2 className="network-title">{NetworkInfo.networkName}</h2>
            <div className="network-list">
                {NetworkInfo.Images.map((container, index) => (
                    <div key={index} className="network-item">
                        <h3 className="network-item-element">
                            {container.Names}
                        </h3>

                        <p className="network-item-element">
                            {container.Status}
                        </p>
                        
                        <p className="network-item-element">
                            {/* {formatPorts(container.Ports)} */}
                            {NetworkInfo.networkName != "host" ? formatPorts(container.Ports) : <></>}
                        </p>
                    </div>
                ))}
            </div>
        </div>
    </>
}

export default Network;