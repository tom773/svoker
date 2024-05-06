type ChipDenominations = {
    [denomination: number]: number;
}
class PlayerType {
    name: string;
    chips: number;
    avatar: string;

    constructor(name: string, chips: number, avatar: string) {
        this.name = name;
        this.chips = chips;
        this.avatar = avatar;
    }

    calculateChipBreakdown(): ChipDenominations {
        const denominations = [1000, 500, 200, 100, 50, 10]; 
        const chips: ChipDenominations = {};
        let remainingBalance = this.chips;
        const weights = {
            1000: 3,
            500: 10,
            200: 9,
            100: 8,
            50: 3,
            10: 1
        };

        // Calculate total weight for normalization
        let totalWeight = 0;
        denominations.forEach(denom => {
            totalWeight += weights[denom];
        });

        // Distribute chips based on weightings
        for (let denom of denominations) {
            // Calculate max chips for current denomination based on its weight
            let maxChips = Math.floor((remainingBalance / denom) * (weights[denom] / totalWeight));
            chips[denom] = maxChips;
            remainingBalance -= maxChips * denom;

            // Adjust totalWeight and weights for the next iteration
            totalWeight -= weights[denom];
            weights[denom] = 0;
        }

        // Handle any remaining balance with the smallest denomination
        if (remainingBalance > 0) {
            let smallestDenom = denominations[denominations.length - 1];
            chips[smallestDenom] += Math.ceil(remainingBalance / smallestDenom);
        }

        return chips;
    }
}

export { PlayerType } ;
