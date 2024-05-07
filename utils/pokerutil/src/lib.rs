#![allow(unused)]
#![allow(dead_code)]
use std::convert::TryFrom;
use wasm_bindgen::prelude::*;
use serde::{Serialize, Deserialize};
use std::collections::HashMap;

#[derive(Debug, Hash, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Serialize)]
pub enum Rank {
    Two = 2,
    Three,
    Four,
    Five,
    Six,
    Seven,
    Eight,
    Nine,
    Ten,
    Jack,
    Queen,
    King,
    Ace
}

#[derive(Hash, Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Serialize)]
pub enum Suit {
    Hearts,
    Diamonds,
    Clubs,
    Spades
}

#[derive(Debug, Hash, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Serialize)]
pub struct Card {
    rank: Rank,
    suit: Suit
}

#[derive(Debug, Hash, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Serialize)]
pub enum Hand {
    HighCard,
    Pair,
    TwoPair,
    ThreeOfAKind,
    Straight,
    Flush,
    FullHouse,
    Quads,
    StraightFlush,
    RoyalFlush
}

#[derive(Debug, Hash, Clone, PartialEq, Eq, PartialOrd, Ord, Serialize)]
pub struct BestHand {
    ranks: Vec<Rank>,
    hand: Hand,
}

pub fn highest_card(card_labels: &mut Vec<String>) -> Vec<Card>{
    // Sort highest
    let mut cards = card_labels.iter().filter_map(|label| hand(&label)).collect::<Vec<Card>>();
    cards.sort_by_key(|label| {
        label.rank
    });
    return cards;
}
#[wasm_bindgen]
pub fn is_pair(mut card_labels: Vec<String>) -> String {
    

    let mut best_hand = BestHand {ranks: Vec::new(), hand: Hand::HighCard};

    let mut cards = card_labels.iter().filter_map(|label| hand(&label)).collect::<Vec<Card>>();
    let mut rank_count = HashMap::new(); 
    
    for card in cards.iter() {
        let count = rank_count.entry(card.rank).or_insert(0);
        *count += 1;
    }
    let pairs = rank_count.iter().filter(|(_, &count)| count == 2).count();
    let three_of_a_kind = rank_count.iter().filter(|(_, &count)| count == 3).count();
    let quads = rank_count.iter().filter(|(_, &count)| count == 4).count();
    
    println!("{:?}", rank_count);

    if (quads != 0) {
        best_hand.hand = Hand::Quads;
        let quads_ranks = rank_count.iter().filter(|(_, &label)| label == 4).map(|(key, _)| *key).collect::<Vec<Rank>>();
        for rank in quads_ranks.iter() {
            if (!best_hand.ranks.contains(rank)) {
                best_hand.ranks.push(*rank);
            }
        }
        println!("Best Hand {:?}", best_hand);
        return serde_json::to_string(&best_hand).unwrap();
    }
    if (three_of_a_kind != 0) {
        if (pairs != 0) {
            best_hand.hand = Hand::FullHouse;
            let three_of_a_kind_ranks = rank_count.iter().filter(|(_, &label)| label == 3).map(|(key, _)| *key).collect::<Vec<Rank>>();
            for rank in three_of_a_kind_ranks.iter() {
                if (!best_hand.ranks.contains(rank)) {
                    best_hand.ranks.push(*rank);
                }
            }
            let paired_ranks = rank_count.iter().filter(|(_, &label)| label == 2).map(|(key, _)| *key).collect::<Vec<Rank>>();
            for rank in paired_ranks.iter() {
                if (!best_hand.ranks.contains(rank)) {
                    best_hand.ranks.push(*rank);
                }
            }
        } else {
            best_hand.hand = Hand::ThreeOfAKind;
            let three_of_a_kind_ranks = rank_count.iter().filter(|(_, &label)| label == 3).map(|(key, _)| *key).collect::<Vec<Rank>>();
            for rank in three_of_a_kind_ranks.iter() {
                if (!best_hand.ranks.contains(rank)) {
                    best_hand.ranks.push(*rank);
                }
            }
        }
    } else {
        for (rank, count) in rank_count.iter() {
            match pairs{
                1 => { best_hand.hand = Hand::Pair;
                    let paired_ranks = rank_count.iter().filter(|(_, &label)| label == 2).map(|(key, _)| *key).collect::<Vec<Rank>>();
                    for rank in paired_ranks.iter() {
                        if (!best_hand.ranks.contains(rank)) {
                            best_hand.ranks.push(*rank);
                        }
                    } 
                }
                2 => { best_hand.hand = Hand::TwoPair;
                    let paired_ranks = rank_count.iter().filter(|(_, &label)| label == 2).map(|(key, _)| *key).collect::<Vec<Rank>>();
                    for rank in paired_ranks.iter() {
                        if (!best_hand.ranks.contains(rank)) {
                            best_hand.ranks.push(*rank);
                        }
                    } 
                }
                _ => { best_hand.hand = Hand::HighCard; best_hand.ranks.push(*rank); }
            }
        }
    }
    if (best_hand.hand != Hand::FullHouse){
        best_hand.ranks.sort_by(|a, b| b.cmp(a));
    }
    println!("Best Hand {:?}", best_hand);
    return serde_json::to_string(&best_hand).unwrap();
}

pub fn hand(card_label: &str) -> Option<Card> {
    
    if card_label.len() != 2 {
        return None;
    }

    let rank_char = card_label.chars().nth(0).unwrap();
    let suit_char = card_label.chars().nth(1).unwrap();

    let rank = match rank_char {
        '2' => Rank::Two,
        '3' => Rank::Three,
        '4' => Rank::Four,
        '5' => Rank::Five,
        '6' => Rank::Six,
        '7' => Rank::Seven,
        '8' => Rank::Eight,
        '9' => Rank::Nine,
        'T' => Rank::Ten,
        'J' => Rank::Jack,
        'Q' => Rank::Queen,
        'K' => Rank::King,
        'A' => Rank::Ace,
        _ => return None
    };
    let suit = match suit_char {
        'H' => Suit::Hearts,
        'D' => Suit::Diamonds,
        'C' => Suit::Clubs,
        'S' => Suit::Spades,
        _ => return None 
    };

    let card = Card { rank, suit };

    return Some(card);    
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn t1() {
        is_pair(vec!["2D".to_string(), "2C".to_string(), "2H".to_string(), "2S".to_string(), "9C".to_string(), "5D".to_string(), "AC".to_string()]);
    }
}
