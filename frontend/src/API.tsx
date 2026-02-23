import axios from 'axios';
import { GetContainerInfoResponse } from './models/HTTPResponse';

export const getContainerInfo = async (): Promise<GetContainerInfoResponse> => {
    try {
        const response = await axios.get(`/be/containerInfo`);
        return response.data as GetContainerInfoResponse;
    } catch (error) {
        console.error('Error fetching container info:', error);
        throw new Error('Failed to fetch container info');
    }
};