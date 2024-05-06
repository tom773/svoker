#![allow(unused)]
#![allow(dead_code)]
use std::convert::TryFrom;
use wasm_bindgen::prelude::*;
use serde::{Serialize, Deserialize};

#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Serialize)]
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

#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Serialize)]
pub enum Suit {
    Hearts,
    Diamonds,
    Clubs,
    Spades
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Serialize)]
pub struct Card {
    rank: Rank,
    suit: Suit
}

#[wasm_bindgen]
pub fn highest_card(card_labels: Vec<String>) -> String{
    let highest = card_labels.iter().filter_map(|label| hand(&label)).max();
    println!("{:?}", serde_json::to_string(&highest).unwrap());
    return serde_json::to_string(&highest).unwrap();
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
        highest_card(vec!["2H".to_string(), "3C".to_string(), "KD".to_string(), "KH".to_string()]);
    }
}
