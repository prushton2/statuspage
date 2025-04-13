import axios from 'axios';

export const getContainerInfo = async (): Promise<string> => {
    try {
        const response = await axios.get(`${import.meta.env.VITE_BACKEND_URL}/containerInfo`);
        return JSON.stringify(response.data);
    } catch (error) {
        console.error('Error fetching container info:', error);
        throw new Error('Failed to fetch container info');
    }
};