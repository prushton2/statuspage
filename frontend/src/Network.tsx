import "./Network.css"
import { NetworkContainers } from "./models/Network";

export const Network = ({NetworkInfo}: {NetworkInfo: NetworkContainers}) => {
    return <>
        <div className="network-container">
            <h2 className="network-title">{NetworkInfo.networkName}</h2>
            <div className="network-list">
                {NetworkInfo.Images.map((container, index) => (
                    <div key={index} className="network-item">
                        <h3 className="network-item-title">{container.Names}</h3>
                        <p className="network-item-status">
                            Status: {container.Status}
                        </p>
                        {
                            NetworkInfo.networkName != "host" ?
                                <p className="network-item-ip">
                                    IP Address: {container.Ports}
                                </p>
                            : <></>
                        }
                    </div>
                ))}
            </div>
        </div>
    </>
}

export default Network;