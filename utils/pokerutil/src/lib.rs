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
    suits: Vec<Suit>,
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

pub fn is_straight(card_labels: &mut Vec<String>) -> bool{
    
    let mut numbers: Vec<i32> = card_labels.iter()
        .map(|s| match s.chars().next().unwrap(){
            '2'..='9' => s.chars().next().unwrap().to_digit(10).unwrap() as i32,
            'T' => 10,
            'J' => 11,
            'Q' => 12,
            'K' => 13,
            'A' => 14,
            _ => 0,
        }).collect();

    numbers.sort_unstable();
    for i in 0..=numbers.len()-5{
        if numbers[i+4] == numbers[i] + 4{
            if (numbers[i+1] == numbers[i] + 1) && (numbers[i+2] == numbers[i] + 2) && (numbers[i+3] == numbers[i] + 3) {
                return true;
            }
        }
    }
    return false;
}

// Monstrosity of a function - consider refactoring
#[wasm_bindgen]
pub fn is_pair(mut card_labels: Vec<String>) -> String {
    

    let mut best_hand = BestHand {ranks: Vec::new(), suits: Vec::new(), hand: Hand::HighCard};
    let mut cards = card_labels.iter().filter_map(|label| hand(&label)).collect::<Vec<Card>>();
    let mut rank_count = HashMap::new(); 
    let mut suit_count = HashMap::new();
    let is_straight_ = is_straight(&mut card_labels);

    for card in cards.iter() {
        let count = rank_count.entry(card.rank).or_insert(0);
        *count += 1;
    }
    for card in cards.iter() {
        let count = suit_count.entry(card.suit).or_insert(0);
        *count += 1;
    }
    
    let flush = suit_count.iter().filter(|(_, &count)| count >= 5).count();
    let pairs = rank_count.iter().filter(|(_, &count)| count == 2).count();
    let three_of_a_kind = rank_count.iter().filter(|(_, &count)| count == 3).count();
    let quads = rank_count.iter().filter(|(_, &count)| count == 4).count();
    
    if (flush != 0){
        if (is_straight_){
            best_hand.hand = Hand::StraightFlush;
            let flush_ranks = suit_count.iter().filter(|(_, &label)| label == 5).map(|(key, _)| *key).collect::<Vec<Suit>>();
            for suit in flush_ranks.iter() {
                if (!best_hand.suits.contains(suit)) {
                    best_hand.suits.push(*suit);
                }
            }
            return serde_json::to_string(&best_hand).unwrap();
        }
        best_hand.hand = Hand::Flush;
        let flush_ranks = suit_count.iter().filter(|(_, &label)| label == 5).map(|(key, _)| *key).collect::<Vec<Suit>>();
        for suit in flush_ranks.iter() {
            if (!best_hand.suits.contains(suit)) {
                best_hand.suits.push(*suit);
            }
        }
        return serde_json::to_string(&best_hand).unwrap();
    }
    if (quads != 0) {
        best_hand.hand = Hand::Quads;
        let quads_ranks = rank_count.iter().filter(|(_, &label)| label == 4).map(|(key, _)| *key).collect::<Vec<Rank>>();
        for rank in quads_ranks.iter() {
            if (!best_hand.ranks.contains(rank)) {
                best_hand.ranks.push(*rank);
            }
        }
        return serde_json::to_string(&best_hand).unwrap();
    }
    if (is_straight_){
        best_hand.hand = Hand::Straight;
        let straight_ranks = rank_count.iter().filter(|(_, &label)| label >= 1).map(|(key, _)| *key).collect::<Vec<Rank>>();
        for rank in straight_ranks.iter() {
            if (!best_hand.ranks.contains(rank)) {
                best_hand.ranks.push(*rank);
            }
        }
        best_hand.ranks.sort_by(|a, b| b.cmp(a));
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

    // Testing straight function
    #[test]
    fn test_staight() {
        let coom = is_straight(&mut vec!["2H".to_string(), "3H".to_string(), "4H".to_string(), "5H".to_string(), "6H".to_string()]);
        assert_eq!(coom, true);
    }
    #[test]
    fn not_straight(){
        let coom = is_straight(&mut vec!["2H".to_string(), "3H".to_string(), "4H".to_string(), "5H".to_string(), "7H".to_string()]);
        assert_eq!(coom, false);
    }
    #[test]
    fn test_seven() {
        let coom = is_straight(&mut vec!["2H".to_string(), "3H".to_string(), "4H".to_string(), "5H".to_string(), "6H".to_string(), "7H".to_string(), "8H".to_string()]);
        assert_eq!(coom, true);
    }
    #[test]
    fn test_unordered(){
        let coom = is_straight(&mut vec!["2H".to_string(), "4H".to_string(), "3H".to_string(), "5H".to_string(), "6H".to_string(), "7H".to_string(), "8H".to_string()]);
    }
    #[test]
    fn test_picture(){
        let coom = is_straight(&mut vec!["TH".to_string(), "JH".to_string(), "QH".to_string(), "KH".to_string(), "AH".to_string()]);
        assert_eq!(coom, true);
    }
    #[test]
    fn test_picture_unordered(){
        let coom = is_straight(&mut vec!["AH".to_string(), "JH".to_string(), "QH".to_string(), "KH".to_string(), "TH".to_string()]);
        assert_eq!(coom, true);
    }
    #[test]
    fn test_picture_seven_unordered(){
        let coom = is_straight(&mut vec!["AH".to_string(), "JH".to_string(), "4H".to_string(), "KH".to_string(), "TH".to_string(), "3H".to_string(), "QD".to_string()]);
        assert_eq!(coom, true);
    }

    // Testing pair function
    // Write some tests. Particularly edge cases eg 3 pairs in 7 cards
    #[test]
    fn test_pair() {
        let coom = is_pair(vec!["2H".to_string(), "2D".to_string(), "9C".to_string(), "4D".to_string(), "5C".to_string(), "6S".to_string(), "7S".to_string()]);
        assert_eq!(coom, "{\"ranks\":[\"Two\"],\"suits\":[],\"hand\":\"Pair\"}");
    }
    #[test]
    fn test_two_pair() {
        let coom = is_pair(vec!["2S".to_string(), "2D".to_string(), "3C".to_string(), "3H".to_string(), "5S".to_string(), "6D".to_string(), "7C".to_string()]);
        assert_eq!(coom, "{\"ranks\":[\"Three\",\"Two\"],\"suits\":[],\"hand\":\"TwoPair\"}");
    }
    #[test]
    fn test_three_of_a_kind() {
        let coom = is_pair(vec!["2H".to_string(), "2D".to_string(), "2C".to_string(), "3S".to_string(), "5S".to_string(), "6D".to_string(), "7C".to_string()]);
        assert_eq!(coom, "{\"ranks\":[\"Two\"],\"suits\":[],\"hand\":\"ThreeOfAKind\"}");
    }
    #[test]
    fn test_full_house() {
        let coom = is_pair(vec!["2H".to_string(), "2D".to_string(), "2C".to_string(), "3D".to_string(), "3H".to_string(), "6H".to_string(), "7H".to_string()]);
        assert_eq!(coom, "{\"ranks\":[\"Two\",\"Three\"],\"suits\":[],\"hand\":\"FullHouse\"}");
    }
    #[test]
    fn test_quads() {
        let coom = is_pair(vec!["2H".to_string(), "2D".to_string(), "2C".to_string(), "2S".to_string(), "3H".to_string(), "6H".to_string(), "7H".to_string()]);
        assert_eq!(coom, "{\"ranks\":[\"Two\"],\"suits\":[],\"hand\":\"Quads\"}");
    }
    #[test]
    fn test_straight() {
        let coom = is_pair(vec!["2H".to_string(), "3D".to_string(), "4C".to_string(), "5S".to_string(), "6H".to_string(), "JH".to_string(), "KH".to_string()]);
        assert_eq!(coom, "{\"ranks\":[\"King\",\"Jack\",\"Six\",\"Five\",\"Four\",\"Three\",\"Two\"],\"suits\":[],\"hand\":\"Straight\"}");
    }
}

