export class Strain {
    id: string = '';
    producer_id : string = '';
    name: string = '';
    category: string = '';
    genetics: string = '';
    thc: number = 0;
    terpenes: number = 0;
    total_cannabinoids: number = 0;
    cbd: number = 0;
    terpene_list: string[] = new Array<string>(3);
    harvest_date: string = '';
    price: number = 0;
    images: string[] = [];
}