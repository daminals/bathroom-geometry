export type Bathroom = {
	id: number;
	name: string;
    gender: 'M' | 'F' | 'U';
    accessible: boolean;
    menstrualProducts: boolean; 
    color?: string;
};