import { Strain } from './strain'

export class OrderItem {
    strain: Strain = new Strain()
    quantity: number = 0
}

export class OrderInformation {
    company_name: string = '';
    license_number: string = '';
    email: string = '';
    phone: string = '';
    address: OrderInformationAddress = new OrderInformationAddress()
}

class OrderInformationAddress {
    street: string = '';
    city: string = '';
    state: string = '';
    postal: string = '';
}