import { axiosHttpInstance } from './axios'

export const RoomServices = {
    getAllRooms() {
        return axiosHttpInstance.get('/rooms/')
    },

    getContactsList(username: string) {
        return axiosHttpInstance.get('/rooms/contact-list', { params: { username: username } })
    },

    getChatHistory(user1: string, user2: string) {
        return axiosHttpInstance.get(`/rooms/chat-history`, { params: { u1: user1, u2: user2 } })
    },
}
