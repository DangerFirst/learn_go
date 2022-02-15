package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//withCancel()
	//withTimeout()
	//withValue()
	withDeadline()
}

func withDeadline() {
	now := time.Now()
	newTime := now.Add(1 * time.Second)
	ctx, _ := context.WithDeadline(context.TODO(), newTime)
	go tv(ctx)
	go mobile(ctx)
	go game(ctx)
	time.Sleep(2 * time.Second)
}

func tv(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关电视")
			return
		default:
		}
		fmt.Println("看电视")
		time.Sleep(300 * time.Millisecond)
	}
}

func mobile(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关手机")
			return
		default:
		}
		fmt.Println("玩手机")
		time.Sleep(300 * time.Millisecond)
	}
}

func game(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关游戏机")
			return
		default:
		}
		fmt.Println("玩游戏机")
		time.Sleep(300 * time.Millisecond)
	}
}

func withValue() {
	ctx := context.WithValue(context.TODO(), "1", "钱包")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("withValue:1", ctx.Value("1"))
		fmt.Println("withValue:2", ctx.Value("2"))
		fmt.Println("withValue:3", ctx.Value("3"))
		fmt.Println("withValue:4", ctx.Value("4"))
	}(ctx)
	goToPapa(ctx)
	time.Sleep(20 * time.Second)
}

func goToPapa(ctx context.Context) {
	ctx = context.WithValue(ctx, "2", "充电宝")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("goToPapa:1", ctx.Value("1"))
		fmt.Println("goToPapa:2", ctx.Value("2"))
		fmt.Println("goToPapa:3", ctx.Value("3"))
		fmt.Println("goToPapa:4", ctx.Value("4"))
	}(ctx)
	goToMama(ctx)
}

func goToMama(ctx context.Context) {
	ctx = context.WithValue(ctx, "3", "小夹克")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("goToMama:1", ctx.Value("1"))
		fmt.Println("goToMama:2", ctx.Value("2"))
		fmt.Println("goToMama:3", ctx.Value("3"))
		fmt.Println("goToMama:4", ctx.Value("4"))
	}(ctx)
	goToGrandma(ctx)
}

func goToGrandma(ctx context.Context) {
	ctx = context.WithValue(ctx, "4", "大苹果")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("goToGrandma:1", ctx.Value("1"))
		fmt.Println("goToGrandma:2", ctx.Value("2"))
		fmt.Println("goToGrandma:3", ctx.Value("3"))
		fmt.Println("goToGrandma:4", ctx.Value("4"))
	}(ctx)
	goToParty(ctx)
}

func goToParty(ctx context.Context) {
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("goToParty:1", ctx.Value("1"))
		fmt.Println("goToParty:2", ctx.Value("2"))
		fmt.Println("goToParty:3", ctx.Value("3"))
		fmt.Println("goToParty:4", ctx.Value("4"))
	}(ctx)

}

func withTimeout() {
	ctx, _ := context.WithTimeout(context.TODO(), 1*time.Second)
	fmt.Println("开始部署望远镜，发送信号")
	go distributeMainFrame(ctx)
	go distributeMainBody(ctx)
	go distributeMainCover(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("任务超时，部署没有完成")
	}
	time.Sleep(20 * time.Second)
}

func distributeMainFrame(ctx context.Context) {
	time.Sleep(10 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务取消：distributeMainFrame")
		return
	default:
	}
	fmt.Println("部署:distributeMainFrame")
}

func distributeMainBody(ctx context.Context) {
	time.Sleep(10 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务取消：distributeMainBody")
		return
	default:
	}
	fmt.Println("部署:distributeMainBody")
}

func distributeMainCover(ctx context.Context) {
	time.Sleep(10 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务取消：distributeMainCover")
		return
	default:
	}
	fmt.Println("部署:distributeMainCover")
}

func withCancel() {
	ctx := context.TODO()
	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("need materials to make egg tart")
	go buyNoodle(ctx)
	go buyOil(ctx)
	go buyEgg(ctx)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("power failure,cancel buy")
	cancel() //当调用cancel后，所有由此上下文衍生出的context都会cancel
	time.Sleep(1 * time.Second)
}

func buyNoodle(ctx context.Context) {
	fmt.Println("go to buy noodle")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("received news,don't buy noodle")
		return
	default:
	}
	fmt.Println("buy noodle")
}

func buyOil(ctx context.Context) {
	fmt.Println("go to buy oil")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("received news,don't buy oil")
		return
	default:
	}
	fmt.Println("buy oil")
}

func buyEgg(ctx context.Context) {
	ctx1, _ := context.WithCancel(ctx)
	fmt.Println("go to buy egg")
	select {
	case <-ctx.Done():
		fmt.Println("received news,don't buy egg")
		return
	default:
	}
	fmt.Println("buy egg")
	go buyAEgg(ctx1)
	go buyBEgg(ctx1)
}

func buyAEgg(ctx context.Context) {
	fmt.Println("go to buy A egg")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("received news,don't buy A egg")
		return
	default:
	}
	fmt.Println("buy A egg")
}

func buyBEgg(ctx context.Context) {
	fmt.Println("go to buy B egg")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("received news,don't buy B egg")
		return
	default:
	}
	fmt.Println("buy B egg")
}
