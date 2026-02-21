import axios from 'axios';
import { Image } from './models/Image';

export const getContainerInfo = async (): Promise<Image[]> => {
    try {
        const response = await axios.get(`/be/containerInfo`);
        return response.data as Image[];
    } catch (error) {
        console.error('Error fetching container info:', error);
        throw new Error('Failed to fetch container info');
    }
};