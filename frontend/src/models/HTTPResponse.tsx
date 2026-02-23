import { Image } from "./Image";

export interface GetContainerInfoResponse {
    containers: Image[],
    topNetworks: string[],
    bottomNetworks: string[]
}