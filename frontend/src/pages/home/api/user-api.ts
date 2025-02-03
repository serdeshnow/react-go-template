import { env } from '@/shared/lib/env.ts';
import axios from 'axios';

const user: { id: number } = {
  id: 1,
};

export const fetchUser = async () => {
  try {
    const response = await axios.get(`${env.API_URL}/user/get/${user.id}`, {
      withCredentials: true,
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
    });

    console.log("Response Status:", response.status);
    console.log("Response Headers:", response.headers);

    return response.data;
  } catch (error: any) {
    console.error("Fetch error:", error?.response?.data || error.message);
    throw new Error(`FETCH_USER Status: ${error?.response?.status || "Unknown"}`);
  }
};
