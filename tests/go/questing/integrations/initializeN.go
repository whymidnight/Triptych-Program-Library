package integrations

import (
	"fmt"
	"log"
	"math"
	"time"

	"triptych.labs/questing"
	"triptych.labs/questing/quests"
	quest_ops "triptych.labs/questing/quests/ops"
	"triptych.labs/utils"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

type StakingQuestScope struct {
	name               string
	duration           int
	left               int
	right              int
	leftCreators       [5]solana.PublicKey
	rightCreators      [5]solana.PublicKey
	yieldPer           int
	yieldPerTime       int
	stakingTokenName   string
	stakingTokenSymbol string
	milestones         *[]questing.Milestone
}

type RewardQuestScope struct {
	name          string
	duration      int
	left          int
	right         int
	leftCreators  [5]solana.PublicKey
	rightCreators [5]solana.PublicKey
	rewards       []questing.Reward
	tender        questing.Tender
	tenderSplits  []questing.Split
}

var GEN1 = [5]solana.PublicKey{
	solana.MustPublicKeyFromBase58("E835GdtAHygnkynfwQkmxPrhfxYfsVi6Kfr7gFx6NmkT"),
	solana.MustPublicKeyFromBase58("E835GdtAHygnkynfwQkmxPrhfxYfsVi6Kfr7gFx6NmkT"),
	solana.MustPublicKeyFromBase58("E835GdtAHygnkynfwQkmxPrhfxYfsVi6Kfr7gFx6NmkT"),
	solana.MustPublicKeyFromBase58("E835GdtAHygnkynfwQkmxPrhfxYfsVi6Kfr7gFx6NmkT"),
	solana.MustPublicKeyFromBase58("E835GdtAHygnkynfwQkmxPrhfxYfsVi6Kfr7gFx6NmkT"),
}

var GEN2 = [5]solana.PublicKey{
	solana.MustPublicKeyFromBase58("E835GdtAHygnkynfwQkmxPrhfxYfsVi6Kfr7gFx6NmkT"),
	solana.MustPublicKeyFromBase58("E835GdtAHygnkynfwQkmxPrhfxYfsVi6Kfr7gFx6NmkT"),
	solana.MustPublicKeyFromBase58("E835GdtAHygnkynfwQkmxPrhfxYfsVi6Kfr7gFx6NmkT"),
	solana.MustPublicKeyFromBase58("E835GdtAHygnkynfwQkmxPrhfxYfsVi6Kfr7gFx6NmkT"),
	solana.MustPublicKeyFromBase58("E835GdtAHygnkynfwQkmxPrhfxYfsVi6Kfr7gFx6NmkT"),
}

func CreateNStakingQuests() {

	scopes := []StakingQuestScope{
		{
			name:               "Pool Zero",
			duration:           0,
			left:               1,
			right:              0,
			leftCreators:       GEN1,
			rightCreators:      GEN2,
			yieldPer:           50,
			yieldPerTime:       10,
			stakingTokenName:   "P0",
			stakingTokenSymbol: "qstPZero",
			milestones: &[]questing.Milestone{
				{
					Tick:     2,
					Modifier: 5,
				},
				{
					Tick:     4,
					Modifier: 10,
				},
				{
					Tick:     8,
					Modifier: 5,
				},
			},
		},
		{
			name:               "Marine Cats",
			duration:           0,
			left:               2,
			right:              0,
			leftCreators:       GEN1,
			rightCreators:      GEN2,
			yieldPer:           50,
			yieldPerTime:       10,
			stakingTokenName:   "Sea Weeds",
			stakingTokenSymbol: "qstSWEED",
			milestones: &[]questing.Milestone{
				{
					Tick:     2,
					Modifier: 5,
				},
				{
					Tick:     4,
					Modifier: 10,
				},
				{
					Tick:     8,
					Modifier: 5,
				},
			},
		},
		{
			name:               "Talking Trees",
			duration:           0,
			left:               4,
			right:              0,
			leftCreators:       GEN1,
			rightCreators:      GEN2,
			yieldPer:           150,
			yieldPerTime:       10,
			stakingTokenName:   "WOOD",
			stakingTokenSymbol: "qstWOOD",
			milestones: &[]questing.Milestone{
				{
					Tick:     2,
					Modifier: 5,
				},
				{
					Tick:     4,
					Modifier: 10,
				},
				{
					Tick:     8,
					Modifier: 5,
				},
			},
		},
	}

	createNStakingQuests(scopes)
}

func CreateNRewardQuests() {
	rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	ixs := make([]solana.Instruction, 0)
	questRewardIx, questRewardMint := quest_ops.RegisterQuestsStakingReward(oracle.PublicKey(), "qstNBA WL", "qstNBAWL", "")
	ixs = append(ixs, questRewardIx)

	utils.SendTx(
		"list",
		ixs,
		append(make([]solana.PrivateKey, 0), oracle, questRewardMint),
		oracle.PublicKey(),
	)

	tenderMint := solana.MustPublicKeyFromBase58("3tQsckZ7R9Tec2rGghGL8CD8RGHTMcoZ8XRQRJnwujGr")

	tenderMintMeta := utils.GetTokenMintData(rpcClient, tenderMint)

	scopes := []RewardQuestScope{
		{
			name:          "60% Whitelist",
			duration:      60 * 60 * 24,
			left:          1,
			right:         1,
			leftCreators:  GEN1,
			rightCreators: GEN2,
			rewards: []questing.Reward{
				{
					MintAddress:   questRewardMint.PublicKey(),
					Threshold:     60,
					Amount:        10,
					AuthorityEnum: 0,
					Cap:           math.MaxInt64,
					Counter:       0,
				},
			},
			tender: questing.Tender{
				MintAddress: tenderMint,
				Amount:      utils.ConvertUiAmountToAmount(float64(10), tenderMintMeta.Decimals),
			},
			tenderSplits: []questing.Split{
				{
					TokenAddress: solana.PublicKey{},
					OpCode:       0,
					Share:        100,
				},
			},
		},
		{
			name:          "100% Whitelist",
			duration:      60 * 60 * 24,
			left:          1,
			right:         1,
			leftCreators:  GEN1,
			rightCreators: GEN2,
			rewards: []questing.Reward{
				{
					MintAddress:   questRewardMint.PublicKey(),
					Threshold:     100,
					Amount:        10,
					AuthorityEnum: 0,
					Cap:           math.MaxInt64,
					Counter:       0,
				},
			},
			tender: questing.Tender{
				MintAddress: tenderMint,
				Amount:      utils.ConvertUiAmountToAmount(float64(20), tenderMintMeta.Decimals),
			},
			tenderSplits: []questing.Split{
				{
					TokenAddress: solana.PublicKey{},
					OpCode:       0,
					Share:        100,
				},
			},
		},
	}

	createNRewardQuests(scopes)
}

func createNStakingQuests(scopes []StakingQuestScope) {
	rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	stakingMints := make([]solana.PrivateKey, 0)

	ixs := make([]solana.Instruction, 0)

	for _, scope := range scopes {
		stakingRewardIx, stakingMint := quest_ops.RegisterQuestsStakingReward(oracle.PublicKey(), scope.stakingTokenName, scope.stakingTokenSymbol, "")
		ixs = append(ixs, stakingRewardIx)
		stakingMints = append(stakingMints, stakingMint)
	}

	signers := make([]solana.PrivateKey, 0)
	signers = append(signers, oracle)
	for _, stakingMint := range stakingMints {
		signers = append(signers, stakingMint)
	}

	utils.SendTx(
		"list",
		ixs,
		signers,
		oracle.PublicKey(),
	)

	log.Println("sleeping...")
	time.Sleep(5 * time.Second)

	questsPda, _ := quests.GetQuests(oracle.PublicKey())
	questsData := quests.GetQuestsData(rpcClient, questsPda)
	for i, scope := range scopes {
		questIxs := make([]solana.Instruction, 0)
		questData := questing.Quest{
			Index:           questsData.Quests + uint64(i),
			Name:            scope.name,
			Duration:        int64(scope.duration),
			Oracle:          oracle.PublicKey(),
			WlCandyMachines: []solana.PublicKey{oracle.PublicKey()},
			Tender:          nil,
			TenderSplits:    nil,
			Rewards:         []questing.Reward{},
			StakingConfig: &questing.StakingConfig{
				MintAddress:  stakingMints[i].PublicKey(),
				YieldPer:     uint64(scope.yieldPer),     // 10 secounds
				YieldPerTime: uint64(scope.yieldPerTime), // 5 tokens
			},
			PairsConfig: &questing.PairsConfig{
				Left:          uint8(scope.left),
				LeftCreators:  scope.leftCreators,
				Right:         uint8(scope.right),
				RightCreators: scope.rightCreators,
			},
			Milestones: scope.milestones,
		}

		creationIx, _ := quest_ops.CreateQuest(rpcClient, oracle.PublicKey(), questData, true)
		questIxs = append(questIxs, creationIx)

		utils.SendTx(
			"list",
			questIxs,
			append(make([]solana.PrivateKey, 0), oracle),
			oracle.PublicKey(),
		)
	}
}

func createNRewardQuests(scopes []RewardQuestScope) {
	// CTM8npagWrtdi85aYix3kpD23yKdboPFMXk9fPWMBoD7
	rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	questsPda, _ := quests.GetQuests(oracle.PublicKey())
	questsData := quests.GetQuestsData(rpcClient, questsPda)
	for i, scope := range scopes {
		questIxs := make([]solana.Instruction, 0)
		questData := questing.Quest{
			Index:           questsData.Quests + uint64(i),
			Name:            scope.name,
			Duration:        int64(scope.duration),
			Oracle:          oracle.PublicKey(),
			WlCandyMachines: []solana.PublicKey{oracle.PublicKey()},
			Tender:          &scope.tender,
			TenderSplits:    &scope.tenderSplits,
			Rewards:         scope.rewards,
			StakingConfig:   nil,
			PairsConfig: &questing.PairsConfig{
				Left:          uint8(scope.left),
				LeftCreators:  scope.leftCreators,
				Right:         uint8(scope.right),
				RightCreators: scope.rightCreators,
			},
		}
		fmt.Println(scope.tender, scope.tenderSplits)

		creationIx, _ := quest_ops.CreateQuest(rpcClient, oracle.PublicKey(), questData)
		questIxs = append(questIxs, creationIx)

		utils.SendTx(
			"list",
			questIxs,
			append(make([]solana.PrivateKey, 0), oracle),
			oracle.PublicKey(),
		)
	}
}
