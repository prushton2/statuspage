import { HealthSummary } from "../models/HealthSummary";
import "./Ribbon.css"

function Ribbon({healthSummary}: {healthSummary: HealthSummary}) {
    
    return (<>
        <div 
            className={`Summary ${healthSummary.unhealthy == 0 ? "Success" : "Fail"}`} 
            title={healthSummary.unhealthy == 0 ? "All Services Running" : `${healthSummary.unhealthy} Service${healthSummary.unhealthy == 1 ? " is" : "s are"} Unhealthy`}>
            {healthSummary.unhealthy == 0 ? "✓" : "X"}
        </div>
        <div className="Ribbon">
            linode.prushton.com
        </div>
    </>);
}

export default Ribbon;