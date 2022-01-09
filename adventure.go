package main

import ("fmt"
		"time"
		"strings")
/*
This is a choose your own adventure type of story where readers
make decisions on how they want the story to be. The story unfolds
differently for different readers and different readers have different
experiences basing on the choices that they make. 
*/
func main(){
adventure()
}
type Scene struct{
	story string
	decisions string
}

func adventure() {
endScene := "-------------------------------------------------------------------------------------"
var input int8
introduction := 
`This is a choose your own adventure type of story where readers
make decisions on how they want the story to be. The story unfolds
differently for different readers and different readers have different
experiences basing on the choices that they make. Readers make choices
by pressing either 1 or 2 and hitting Enter.`
end := "End of story."

sceneA := Scene{
`It is a cold but vibrant Saturday evening. Cold typically because it had rained heavily that afternoon and the weather had remained
rather gloomy. I am taking a stroll with Eve along the university way just outside the school. I had met Eve a year ago in school; 
since that very first day she had grown to be a very important person in my life, in short, we were in a relationship. We then walk 
past a man selling a local evening snack, some roasted meat stuffed in goat intestines, popularly known as ‘mtura’. I had eaten mtura 
ever since I was a kid and had grown to like it. I asked Eve whether she had some coins in her pocket. She didnt even answer, she 
just removed one of her hands from her large hoodie pocket where they were firmly tucked and gave me a sh.50 note and told me that 
she’ll head straight to her room and that I’ll find her there.
`,`1.Let her go. 
2.Convince her to eat with you.`} 
sceneAA := Scene{
`I  lean in towards her, hug her halfway with one hand, then I pecked her neck and whispered in her ear, “Thank you.” She then walks
away wiggling her behind intentionally since she knew I’d look. I bought mtura worth sh.30. I particularly loved that spot since the
mtura guy offered splendid services. He had some local vegetable salad that he garnished with ghost peppers, popularly known as 
‘kachumbari’. He also gave some bonus pieces if a customer had purchased mtura worth sh.30. So I ate it, paid and headed straight to
Eve’s room. Surprisingly, the lights in her room were off, I knocked faintly on the door, then reached out beneath the door hinge to
check whether there was a padlock and indeed it was there.
`,`1.Call her. 
2.Go to my friend’s place.`}
sceneAAA := Scene{
`I quickly took out my phone and dialed her number. She hung up, which was highly unlikely of her. I called her again and this time round
it did not go through. I called Shanice,  who was her close friend and asked her whether Eve was with her. She told me that the last time
they had talked was about three hours ago. I hung up and left the apartment. I had started to panic. I could feel my heart start to beat
faster. I knew it was only a matter of time before her mom calls her only to not find her then call me in order to speak to her. 
`,`1.Check all places that she may be
2.Go to my friends place`}
sceneAAAA := Scene{
`I decided to check all places that I assumed she may have gone. I went to the next apartment where her friend Jane, a notorious gossiper,
lived. I particularly disliked Jane. She was too noisy and had poor turn-taking skills during a conversation. I found Jane but Eve was not
there. I went to the salon, where she usually does her hair, she also wasn’t there. “Had she been kidnapped? No, maybe she has lost her 
phone. Yes, she has most likely lost her phone. But why has she not called me with a different number to explain? Why is she not at her 
place? She told me she’ll be there,” I thought to myself.
`,`1.Continue searching for her
2.Go home`} 
sceneAAAAA := Scene{
`I wasn’t going to my place until I was sure that Eve was not in any form of danger. I decided to trace my steps back to the mtura guy. 
After reaching the spot, I started to contemplate on different scenarios in which she may or may have not followed different paths. I was
literally confused. In the depth of my confusion, a car came out of nowhere, I was grabbed by the neck by some guy that I didn’t see and
was thrown in the car. I only remember that it was a black sedan. Inside the car I was gagged using some dirty and smelly rag that made
me throw up and made me really have it rough and the entire mysterious journey even more miserable. I then passed out…
`,`1.Forced to join a gang
2.Enslaved in a foreign country`}
sceneAB := Scene{
`“But what do you want to go and do in your room without me?” I asked her. She just smiled broadly, typical of her and asked me to buy the
mtura. I particularly loved that spot since the mtura guy offered splendid services. He had some local vegetable salad that he garnished 
with ghost peppers, popularly known as ‘kachumbari’. He also gave some bonus pieces if a customer had purchased mtura worth sh.30. So I 
bought mtura worth sh.40 and was given two extra pieces. We then ate, paid and left. We then walked to her place as she had seriously 
suggested.
`,`1.Sleep over at her place 
2.make an excuse and go to my friends place`}
sceneABA := Scene{
`It was around 7:00pm and we were at her place when she asked me what we shall prepare for our supper. I told her that I had a sh.100 note 
and would go to buy chapatis in the next hour. “You just know how insecure this place is, especially during the night. So if it’s buying 
something, you better go as early as possible so that you can come back indoors as soon as possible,” she told me. She was overly paranoid
about the insecurity issue but I didn’t blame her. After all, the country was recovering from from a very gruesome war that it had had two 
years ago against Somalia, particularly against a militia group known as Al Shabaab. The general notion was that our country had won. So
I picked a recyclable carrier bag and left.
`,`1.Go to my friends place first
2.Go buy chapati.`} 
sceneABAA := Scene{
`The best chapati spot I knew was a 15 minute walk away, so it would take me roughly 40 minutes to be back at Eve’s place. I decided to go
to my friend’s place first, play a match or two then go for the chapatis. I called Ray, who lived just two blocks away and asked him whether
I may come over to play a video game football match against him. He said yes and went over feeling very sly. I played two games in which I
lost both then realized that I was running out of time. I left hurriedly and walked fast toward the chapati spot. I bought six of them
then realized that I had some time to spare before going back at Eve’s place. 
`,`1.Go for mtura
2.Head back to Eve’s place`}
sceneABAAA := Scene{
`I had sh 40 left so I decided to go for some mtura at my usual spot, the one I had been at with Eve about an hour ago.  “Is she your
girlfriend?” The mtura guy asked me as he cut mtura worth sh 15 as I had told him. “Yes,” I replied. “But I’ve never seen you with her 
though,” he said and began to laugh. He had asked that since it was the first time I had brought her at that spot. So I told him the same 
and started to feast on my meal. “Once you realize that she is the one, keep her, very few of them exist,” he said as he cut two more pieces
and gave them to me. “Thank you so much,” I replied. I felt that the world needed more of such people and less of the ones that were 
responsible for the war that had claimed thousands of innocent lives. I then paid and left. As I was walking towards the apartment where 
Eve stayed,  a car came out of nowhere, I was grabbed by the neck by some guy that I didn’t see and was thrown in the car. I only remember 
that it was a black sedan. Inside the car I was gagged using some dirty and smelly rag that made me throw up and made me really have it rough
and the entire mysterious journey even more miserable. I then passed out…
`,`1.Forced to join a gang
2.Enslaved in a foreign country`}
sceneAAAB := Scene{
`I was literally confused. I decided to go to my place to contemplate the whole matter and make the right decisions. My place was was just
in the apartment, next to the one in which Eve lived. As I was walking towards my place, I saw Ray and called him. He was carrying something
in a brown, khaki carrier bag.  I didn’t want him to know that Eve was missing so I just told him that I was going home and that I had 
nothing else to do. He told me that he had just bought a bottle of vodka and was going to his place, he even asked me to accompany him so 
that we may take shots and even play video games. I knew I only had a few second to respond without letting him know that Eve may be in danger.
`,`1.Accept Ray’s offer
2.Decline and go to my place.`}
sceneAAABA := Scene{
`“Lets go,” I told Ray, “but I wont stay for long.” he said okay and we headed over toward his place. I Had decided to go to Ray’s place, 
then leave later on, hoping that Eve may have come back from God knows where. At the same time, I strongly felt the urge to tell Ray about 
Eve. We then reached his house and took 3 shots each and played two games in which I lost both. Ray may have noticed the fact that I acted 
withdrawn and asked me what the issue was. I couldn’t help it. I let it all out. “Your wife has disappeared and yet you’re here taking shots 
of vodka as if it’s no big deal. Get up and let’s go look for her!”he said angrily. I felt blissful with a pinch of guilt though. He asked me 
that we go back to the mtura spot that we had previously been. As we walked towards the place, he kept on mocking me on why I was afraid of 
informing him yet someone was in danger. Suddenly, we saw a black sedan car packed a few meters before the mtura spot. It was ray that had 
noticed, this was so because he had a thing for Audi cars, and he just couldn’t stop staring. About two hours later, our efforts became futile
and we decided to report the matter to the police. An unsaved number called me and to our relief, it was Eve. My lovely Eve..
`,``}
sceneAAABB := Scene{
`“Today I feel very worn out bro. I have had a very busy day and I don’t think I’ll make it. But spare some vodka for me for tomorrow. Right
now I just need a good night sleep,” I told him. His frantic efforts to convince me go with him to his place literally fell on deaf ears. 
He had no option but to accept, though half-heartedly. I got into my room and got several panic attacks. “Where exactly would she be. Should
I go to the police? Should I call Ray and tell him about Eve’s disappearance?” Random thoughts tinkered through my mind as a wave of paranoia
swept through me. I fidgeted for about three hours without making up my mind to something constructive. In that miserable melee, an unsaved
number called me and there she was, my Eve. My precious Eve.
`,``}
sceneAAB := Scene{
`I called Ray, who lived just two blocks away and asked him whether I may come over to play a video game football match against him. He said
 yes and went over feeling very confident. I assumed that Eve had gone to see her friend Jane and that she would let me know later on.  I 
 particularly disliked Jane. She was too noisy and had poor turn-taking skills during a conversation. Hence, I did not call her because I knew
  that Jane would gladly demand to talk to me too. I found that Ray had bought a bottle of vodka which we took 3 shots each before going on to 
  play subsequent matches in which he won majority of them. We played until 5:00am then fell asleep.
`,``}


	fmt.Println(endScene)
	t := time.Now()
	switch {
	case t.Hour() < 12:
		nicePrint("Good morning. ")
	case t.Hour() < 17:
		nicePrint("Good afternoon. ")
	default:
		nicePrint("Good evening. ")
	}
	nicePrint(introduction)
	fmt.Println("\n")
	for i:= 1;i<4;i++{
		fmt.Println(endScene)
	}
	fmt.Println("\n")
	nicePrint(sceneA.story)
	fmt.Println(endScene)
	fmt.Println(sceneA.decisions)
	fmt.Scanln(&input)
	if input == 1{
		nicePrint(sceneAA.story)
		fmt.Println(endScene)
		fmt.Println(sceneAA.decisions)
		fmt.Scanln(&input)
		if input == 1{
			nicePrint(sceneAAA.story)
			fmt.Println(endScene)
			fmt.Println(sceneAAA.decisions)
			fmt.Scanln(&input)
			if input == 1{
				nicePrint(sceneAAAA.story)
				fmt.Println(endScene)
				fmt.Println(sceneAAAA.decisions)
				fmt.Scanln(&input)
				if input == 1{
					nicePrint(sceneAAAAA.story)
					fmt.Println(endScene)
					fmt.Println(sceneAAAAA.decisions)
					fmt.Scanln(&input)
				} else if input == 2{
					nicePrint(sceneAAAB.story)
					fmt.Println(endScene)
					fmt.Println(sceneAAAB.decisions)
					fmt.Scanln(&input)
					if input == 1{
						nicePrint(sceneAAABA.story)
						fmt.Println(endScene)
						fmt.Println(sceneAAABA.decisions)
						fmt.Scanln(&input)
					} else if input == 2{
						nicePrint(sceneAAABB.story)
						fmt.Println(endScene)
						fmt.Println(sceneAAABB.decisions)
						fmt.Scanln(&input)
					} else{
						fmt.Println("Invalid choice! Enter either 1 or 2")
					}
				} else{
				fmt.Println("Invalid choice! Enter either 1 or 2")
				}
			} else if input == 2{
				nicePrint(sceneAAAB.story)
				fmt.Println(endScene)
				fmt.Println(sceneAAAB.decisions)
				fmt.Scanln(&input)
				if input == 1{
					nicePrint(sceneAAABA.story)
					fmt.Println(endScene)
					fmt.Println(sceneAAABA.decisions)
					fmt.Scanln(&input)
				} else if input == 2{
					nicePrint(sceneAAABB.story)
					fmt.Println(endScene)
					fmt.Println(sceneAAABB.decisions)
					fmt.Scanln(&input)
				} else{
					fmt.Println("Invalid choice! Enter either 1 or 2")
				}
			} else{
				fmt.Println("Invalid choice! Enter either 1 or 2")
			}
		} else if input == 2{
			nicePrint(sceneAAB.story)
			fmt.Println(endScene)
			fmt.Println(sceneAAB.decisions)
					
			fmt.Scanln(&input)
		} else{
			fmt.Println("Invalid choice! Enter either 1 or 2")
		}
	} else if input == 2{
		nicePrint(sceneAB.story)
		fmt.Println(endScene)
		fmt.Println(sceneAB.decisions)
		fmt.Scanln(&input)
		if input == 1{
			nicePrint(sceneABA.story)
			fmt.Println(endScene)
			fmt.Println(sceneABA.decisions)
			fmt.Scanln(&input)
			if input == 1{
				nicePrint(sceneABAA.story)
				fmt.Println(endScene)
				fmt.Println(sceneABAA.decisions)
				fmt.Scanln(&input)
				if input == 1{
					nicePrint(sceneABAAA.story)
					fmt.Println(endScene)
					fmt.Println(sceneABAAA.decisions)
					fmt.Scanln(&input)
				} else if input == 2{
					nicePrint(sceneA.story)
					fmt.Println(endScene)
					fmt.Println(sceneA.decisions)
					fmt.Scanln(&input)
				} else{
					fmt.Println("Invalid choice! Enter either 1 or 2")
				}
			} else if input == 2{
				nicePrint(sceneA.story)
				fmt.Println(endScene)
				fmt.Println(sceneA.decisions)
				fmt.Scanln(&input)
			} else{
				fmt.Println("Invalid choice! Enter either 1 or 2")
			}
		} else if input == 2{
			fmt.Scanln(&input)
		} else{
			fmt.Println("Invalid choice! Enter either 1 or 2")
		}
		
	} else{
		fmt.Println("Invalid choice! Enter either 1 or 2")
	}
	fmt.Println(end)
	wordPrint("WRITTEN BY LIH")
	nicePrint("13th March 2020")
}
func nicePrint(a string) {
	for _, c := range a{
		fmt.Printf(string(c))
		time.Sleep(30*time.Millisecond)
	}
}
func wordPrint(a string){
	words := strings.Fields(a)
	for _, word:= range words{
		fmt.Println(word)
		time.Sleep(1*time.Second)
	}
}
