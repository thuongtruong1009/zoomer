import axiosInstance from './axios'

export const RoomServices = {
    getAllRooms() {
        return axiosInstance.get('/rooms/')
    },

    getContactsList(username: string) {
        return axiosInstance.get('/rooms/contact-list', { params: { username: username } })
    },

    getChatHistory(user1: string, user2: string) {
        return axiosInstance.get(`/rooms/chat-history`, { params: { u1: user1, u2: user2 } })
    },
}
