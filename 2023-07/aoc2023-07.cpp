#include <fstream>
#include <iostream>
#include <string>
#include <unordered_map>
#include <regex>
#include <algorithm>

using namespace std;

enum class HandType
{
    HighCard,
    OnePair,
    TwoPair,
    ThreeOfAKind,
    FullHouse,
    FourOfAKind,
    FiveOfAKind
};

class Hand
{
protected:
    string cards;
    int bid;
    HandType handType;

    HandType calculateHandType(string cards)
    {
        std::map<char, int> labelCounts;
        for (const auto &card : cards)
        {
            labelCounts[card]++;
        }

        int pairs = 0;
        int threes = 0;
        int fours = 0;
        int fives = 0;

        for (const auto &count : labelCounts)
        {
            if (count.second == 2)
                pairs++;
            else if (count.second == 3)
                threes++;
            else if (count.second == 4)
                fours = 1;
            else if (count.second == 5)
                fives = 1;
        }

        if (fives)
        {
            return HandType::FiveOfAKind;
        }
        else if (fours)
        {
            return HandType::FourOfAKind;
        }
        else if (threes && pairs)
        {
            return HandType::FullHouse;
        }
        else if (threes)
        {
            return HandType::ThreeOfAKind;
        }
        else if (pairs == 2)
        {
            return HandType::TwoPair;
        }
        else if (pairs == 1)
        {
            return HandType::OnePair;
        }
        else
        {
            return HandType::HighCard;
        }
    }

    virtual void assignHandType()
    {
        handType = calculateHandType(cards);
    }

public:
    Hand(string cardsVal, int bidVal) : cards(cardsVal), bid(bidVal)
    {
        assignHandType();
    }

    string getCards() const { return cards; }
    int getBid() const { return bid; }
    HandType getHandType() const { return handType; }

    std::string toString() const
    {
        ostringstream oss;
        oss << "cards:" << cards << " bid: " << bid << " type: " << static_cast<int>(handType);
        return oss.str();
    }
};

class JokerHand : public Hand
{
private:
    string replacementCards = "23456789TQKA";

    void getMaximumJokerType(string cards, int depth)
    {
        if (handType >= HandType::FiveOfAKind)
        {
            return;
        }
        if (depth > 4)
        {
            return;
        }

        for (int i = depth; i < 5; i++)
        {
            if (cards[i] != 'J')
            {
                continue;
            }
            for (const char newCard : replacementCards)
            {
                cards[i] = newCard;
                auto newHandsType = calculateHandType(cards);
                if (newHandsType > handType)
                {
                    handType = newHandsType;
                }
                if (handType >= HandType::FiveOfAKind)
                {
                    return;
                }
                getMaximumJokerType(cards, i + 1);
            }
        }
    }

    virtual void assignHandType()
    {
        handType = calculateHandType(cards);
        string cardsCopy = cards;
        getMaximumJokerType(cardsCopy, 0);
    }

public:
    JokerHand(string cardsVal, int bidVal) : Hand(cardsVal, bidVal)
    {
        assignHandType();
    }
};

bool compareHands(const Hand &lhs, const Hand &rhs, map<char, int> cardWeights)
{
    if (lhs.getHandType() > rhs.getHandType())
    {
        return false;
    }
    else if (lhs.getHandType() < rhs.getHandType())
    {
        return true;
    }
    else
    {
        for (int i = 0; i < lhs.getCards().size(); i++)
        {
            if (cardWeights[lhs.getCards()[i]] > cardWeights[rhs.getCards()[i]])
            {
                return false;
            }
            else if (cardWeights[lhs.getCards()[i]] < cardWeights[rhs.getCards()[i]])
            {
                return true;
            }
        }
        return false;
    }
}

struct HandComparator
{
    bool operator()(const Hand &lhs, const Hand &rhs) const
    {
        map<char, int> cardWeights = {
            {'A', 14},
            {'K', 13},
            {'Q', 12},
            {'J', 11},
            {'T', 10},
            {'9', 9},
            {'8', 8},
            {'7', 7},
            {'6', 6},
            {'5', 5},
            {'4', 4},
            {'3', 3},
            {'2', 2}};

        return compareHands(lhs, rhs, cardWeights);
    }
};

struct JokerHandComparator
{
    bool operator()(const Hand &lhs, const Hand &rhs) const
    {
        map<char, int> cardWeights = {
            {'A', 14},
            {'K', 13},
            {'Q', 12},
            {'T', 10},
            {'9', 9},
            {'8', 8},
            {'7', 7},
            {'6', 6},
            {'5', 5},
            {'4', 4},
            {'3', 3},
            {'2', 2},
            {'J', 1},
        };

        return compareHands(lhs, rhs, cardWeights);
    }
};

Hand getHand(string cards, int bid)
{
    return Hand(cards, bid);
}

Hand getJokerHand(string cards, int bid)
{
    return JokerHand(cards, bid);
}

vector<Hand> load_data(function<Hand(const std::string &, int)> getHand)
{

    vector<Hand> hands;
    regex pattern(R"((\w{5}) (\d+))");

    ifstream file("input.txt");
    if (!file.is_open())
    {
        cerr << "Error opening file" << endl;
        return vector<Hand>();
    }

    string line;
    while (getline(file, line))
    {
        smatch matches;
        if (!regex_search(line, matches, pattern))
        {
            return vector<Hand>();
        }
        hands.push_back(getHand(matches[1], stoi(matches[2])));
    }

    file.close();
    return hands;
}

int solve(vector<Hand> hands)
{

    int rank = 1;
    int total_winnings = 0;
    for (auto hand : hands)
    {
        total_winnings += hand.getBid() * rank;
        // cout << hand.toString() << " rank: " << rank << " total winnings: " << total_winnings << endl;
        rank++;
    };
    return total_winnings;
}

int main()
{
    auto hands = load_data(getHand);
    sort(hands.begin(), hands.end(), HandComparator());
    auto total_winnings = solve(hands);
    cout << "Puzzle 1 - total winnings: " << total_winnings << endl;

    hands = load_data(getJokerHand);
    sort(hands.begin(), hands.end(), JokerHandComparator());
    total_winnings = solve(hands);
    cout << "Puzzle 2 - total winnings: " << total_winnings << endl;

    return 0;
}
