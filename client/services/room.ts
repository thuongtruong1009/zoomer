import { axiosHttpInstance } from './http'

export const RoomServices = {
    getAllRooms() {
        return axiosHttpInstance.get('/rooms/')
    },

    getContactsList(username: string) {
        return axiosHttpInstance.get('/rooms/contact-list', { params: { username: username } })
    },

    getChatHistory(payload: any) {
        return axiosHttpInstance.get(`/rooms/chat-history`, payload)
    },
}
