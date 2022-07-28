# llegend
  - shows segments you are local legend of
  - color coding showing how close to being overtaken
  - color red = someone is nearly passed efforts
  - color green = many efforts before someone catches you
  - color gray = lost LL
  - later: another screen showing ones you are getting close to

## domain:
- llegend.io
- llegend.co
- llegends.io
- llegend.com (2500)


## login
initially, we won't do user login, so everything will operate on the user
with the developer token, rather than the "logged_in" user. Eventually we'll
need to do oauth and login with the user's strava account to make this useful.

## cron job
In order to save on traffic, the LL segments should be loaded once per session
and then periodically the ones which people are LL of should be queried in
a cron job. If the histogram is getting within some threshold an alert can
be sent out.

## requests:
curl: https://www.strava.com/api/v3/segments/19209251/local_legend
which will give you something like this if you are the LL (see histogram):
```
[
    {
        "category": "overall",
        "analytics_context": {
            "segment_id": 19209251,
            "leaderboard_filter_type": "overall",
            "local_legend_athlete_id": 14141827,
            "local_legend_effort_count": 4,
            "viewing_athlete_effort_count": 5,
            "viewing_athlete_bucket": 4,
            "total_athlete_count": 28,
            "total_effort_count": 47,
            "total_buckets": 5
        },
        "window_text": "In the last 90 days",
        "segment": {
            "id": 19209251,
            "name": "Dogleg to Bridge Parkway",
            "activity_type": "Run",
            "distance": 971.6,
            "average_grade": 0.0,
            "elevation_high": 4.0,
            "elevation_low": 3.0,
            "map": {
                "polyline": "elcdF~wciVM@[V]Pc@Ns@NYPI?g@`@MBGD_Av@WLc@ZYHONG??G`@Gd@UTQb@U\\[VONOPIVQn@WDEXKLGTA\\@PK\\CFEDIN@XMh@KXDDDJNF`@FN`@`@FBPFp@IV@v@T`@TRRf@r@d@X\\Zj@`@\\d@N^JLDG?S",
                "lat_lng": [
                    [
                        37.54195,
                        -122.24912
                    ],
                    [
                        37.54202,
                        -122.24913000000001
                    ],
                    [
                        37.54216,
                        -122.24925
                    ],
                    [
                        37.54231,
                        -122.24934
                    ],
                    [
                        37.54249,
                        -122.24942000000001
                    ],
                    [
                        37.542750000000005,
                        -122.24950000000001
                    ],
                    [
                        37.542880000000004,
                        -122.24959000000001
                    ],
                    [
                        37.542930000000005,
                        -122.24959000000001
                    ],
                    [
                        37.543130000000005,
                        -122.24976000000001
                    ],
                    [
                        37.543200000000006,
                        -122.24978000000002
                    ],
                    [
                        37.543240000000004,
                        -122.24981000000001
                    ],
                    [
                        37.54356000000001,
                        -122.25009000000001
                    ],
                    [
                        37.54368,
                        -122.25016000000001
                    ],
                    [
                        37.54386,
                        -122.25030000000001
                    ],
                    [
                        37.54399,
                        -122.25035000000001
                    ],
                    [
                        37.544070000000005,
                        -122.25043000000001
                    ],
                    [
                        37.54411,
                        -122.25043000000001
                    ],
                    [
                        37.54411,
                        -122.25039000000001
                    ],
                    [
                        37.543940000000006,
                        -122.25035000000001
                    ],
                    [
                        37.54375,
                        -122.25024
                    ],
                    [
                        37.54364,
                        -122.25015
                    ],
                    [
                        37.54346,
                        -122.25004000000001
                    ],
                    [
                        37.543310000000005,
                        -122.24990000000001
                    ],
                    [
                        37.54319,
                        -122.24982000000001
                    ],
                    [
                        37.543110000000006,
                        -122.24974000000002
                    ],
                    [
                        37.543020000000006,
                        -122.24969000000002
                    ],
                    [
                        37.5429,
                        -122.24960000000002
                    ],
                    [
                        37.542660000000005,
                        -122.24948
                    ],
                    [
                        37.54263,
                        -122.24945000000001
                    ],
                    [
                        37.542500000000004,
                        -122.24939
                    ],
                    [
                        37.54243,
                        -122.24935
                    ],
                    [
                        37.542320000000004,
                        -122.24934
                    ],
                    [
                        37.542170000000006,
                        -122.24935
                    ],
                    [
                        37.542080000000006,
                        -122.24929000000002
                    ],
                    [
                        37.54193,
                        -122.24927000000001
                    ],
                    [
                        37.54189,
                        -122.24924000000001
                    ],
                    [
                        37.54186,
                        -122.24919000000001
                    ],
                    [
                        37.54178,
                        -122.24920000000002
                    ],
                    [
                        37.541650000000004,
                        -122.24913000000001
                    ],
                    [
                        37.54144,
                        -122.24907
                    ],
                    [
                        37.54131,
                        -122.24910000000001
                    ],
                    [
                        37.54128,
                        -122.24913000000001
                    ],
                    [
                        37.54122,
                        -122.24921
                    ],
                    [
                        37.541180000000004,
                        -122.24938000000002
                    ],
                    [
                        37.541140000000006,
                        -122.24946000000001
                    ],
                    [
                        37.54097,
                        -122.24963000000001
                    ],
                    [
                        37.54093,
                        -122.24965000000002
                    ],
                    [
                        37.54084,
                        -122.24969000000002
                    ],
                    [
                        37.54059,
                        -122.24964000000001
                    ],
                    [
                        37.540470000000006,
                        -122.24965000000002
                    ],
                    [
                        37.54019,
                        -122.24976000000001
                    ],
                    [
                        37.540020000000005,
                        -122.24987000000002
                    ],
                    [
                        37.53992,
                        -122.24997
                    ],
                    [
                        37.53972,
                        -122.25023000000002
                    ],
                    [
                        37.539530000000006,
                        -122.25036000000001
                    ],
                    [
                        37.53938,
                        -122.25050000000002
                    ],
                    [
                        37.53916,
                        -122.25067000000001
                    ],
                    [
                        37.539010000000005,
                        -122.25086
                    ],
                    [
                        37.53893,
                        -122.25102000000001
                    ],
                    [
                        37.53887,
                        -122.25109
                    ],
                    [
                        37.53884,
                        -122.25105
                    ],
                    [
                        37.53884,
                        -122.25095
                    ]
                ]
            }
        },
        "elevation_profile": "https://d3o5xota0a1fcr.cloudfront.net/v6/charts/4YRD65KS5VYJVKLNSVEA2XZ42NL5R5B5G3F7TP7I4TI3XYL55TF5HVBHX63OBR33WGWZIB2BN5JJBCM272F6O===",
        "static_maps": {
            "1x": "https://d3o5xota0a1fcr.cloudfront.net/v6/maps/VQTPR3L3IM3KU4O4GDGO4ZTIBHIDEK5FUKBMPTMT25ATBDLEP7PQRWDZCJKQFSKFXSS3KYNY7URJRCYARZSANW2L4VJA====",
            "2x": "https://d3o5xota0a1fcr.cloudfront.net/v6/maps/INQBSIX6OFJDSTNIMJ2BRGHYRW4QW5LIXD3FPQGKRQJDXTTX4KTIDZ6JPK57ZP6KQEWWUVHPYICZMRQH444LFBX7F26Q===="
        },
        "histogram": {
            "x_labels": [
                {
                    "label": "1",
                    "bucket_index": 0
                },
                {
                    "label": "5",
                    "bucket_index": 4
                }
            ],
            "y_labels": [
                {
                    "label": "1",
                    "position": 0.43,
                    "draw_line": true
                },
                {
                    "label": "10",
                    "position": 0.86,
                    "draw_line": true
                },
                {
                    "label": "20",
                    "position": 1.0,
                    "draw_line": true
                }
            ],
            "athlete_buckets": [
                {
                    "athlete_id": 2021127,
                    "bucket_index": 4
                },
                {
                    "athlete_id": 14141827,
                    "bucket_index": 3
                }
            ],
            "bucket_values": [
                0.98,
                0.69,
                0.69,
                0.43,
                0.43
            ],
            "buckets": [
                {
                    "value": 0.98,
                    "efforts_text": {
                        "text": "18 athletes completed 1 effort",
                        "emphasis": [
                            {
                                "start_index": 0,
                                "length": 11
                            },
                            {
                                "start_index": 22,
                                "length": 8
                            }
                        ]
                    }
                },
                {
                    "value": 0.69,
                    "efforts_text": {
                        "text": "4 athletes completed 2 efforts",
                        "emphasis": [
                            {
                                "start_index": 0,
                                "length": 10
                            },
                            {
                                "start_index": 21,
                                "length": 9
                            }
                        ]
                    }
                },
                {
                    "value": 0.69,
                    "efforts_text": {
                        "text": "4 athletes completed 3 efforts",
                        "emphasis": [
                            {
                                "start_index": 0,
                                "length": 10
                            },
                            {
                                "start_index": 21,
                                "length": 9
                            }
                        ]
                    }
                },
                {
                    "value": 0.43,
                    "efforts_text": {
                        "text": "The Local Legend has done 4 efforts",
                        "emphasis": [
                            {
                                "start_index": 26,
                                "length": 9
                            }
                        ]
                    }
                },
                {
                    "value": 0.43,
                    "efforts_text": {
                        "text": "You completed 5 efforts",
                        "emphasis": [
                            {
                                "start_index": 14,
                                "length": 9
                            }
                        ]
                    }
                }
            ],
            "footer": {
                "icon": "actions_lock_closed_normal",
                "icon_color": "#6d6d78",
                "text": "The Local Legend’s name and efforts are visible to everyone. If you’re not the LCL, your effort count and histogram placement are only visible to you. Tap to learn more.",
                "destination": "strava://support/articles/360043099552"
            }
        },
        "local_legend": {
            "athlete_id": 14141827,
            "title": "Manoj Nayak",
            "profile": "https://d3nn82uaxijpm6.cloudfront.net/assets/avatar/athlete/large-800a7033cc92b2a5548399e26b1ef42414dd1a9cb13b99454222d38d58fd28ef.png",
            "badge_type_id": 1,
            "effort_description": "4 efforts in the last 90 days",
            "mayor_effort_count": 4,
            "show_see_your_results": true,
            "your_efforts_text": {
                "text": "You're only 0 efforts away from becoming the new Local Legend!",
                "emphasis": [
                    {
                        "start_index": 12,
                        "length": 9
                    }
                ]
            }
        },
        "overall_efforts": {
            "total_athletes": "28",
            "total_efforts": "47",
            "total_distance": "45.6 km"
        },
        "your_efforts": {
            "effort_description": "In the last 90 days",
            "total_efforts": "5",
            "total_distance": "4.8 km"
        },
        "leaderboards": [
            {
                "type": "mutual_followers",
                "overall_efforts": {
                    "total_athletes": "1",
                    "total_efforts": "5",
                    "total_distance": "4.8 km"
                },
                "entries": [
                    {
                        "rank": 1,
                        "athlete_id": 2021127,
                        "name": "Jason Ernst",
                        "profile": "https://dgalywyr863hv.cloudfront.net/pictures/athletes/2021127/603244/4/medium.jpg",
                        "badge_type_id": 1,
                        "effort_count": 5,
                        "is_local_legend": false,
                        "destination": "strava://athletes/2021127",
                        "last_effort_text": "1 hour ago"
                    }
                ],
                "footer": {
                    "icon": "navigation_group_normal",
                    "icon_color": "#6d6d78",
                    "text": "This leaderboard ranks athletes you follow who follow you back. To avoid showing up on your followers’ leaderboards you can change your activity privacy to ‘Only You.’ Tap to learn more.",
                    "destination": "strava://support/articles/360043099552"
                },
                "analytics_context": {
                    "top_follower_athlete_id": null,
                    "top_follower_effort_count": null,
                    "following_athlete_count": 1,
                    "following_effort_count": 5
                }
            }
        ],
        "opt_in": {
            "athlete_opted_in": true,
            "text": "The Local Legend's name and efforts are visible to everyone. If you're not the Local Legend (LCL), your effort count and histogram placement are only visible to you. Tap to learn more.",
            "icon": "navigation_group_normal",
            "icon_color": "#6d6d78",
            "learn_more": {
                "title": "Learn More",
                "rows": [
                    {
                        "title": "Learn About Local Legends",
                        "destination": "strava://support/articles/360043099552"
                    }
                ]
            },
            "privacy": {
                "title": "Privacy Control",
                "button_text": "Leave Local Legends",
                "body": "If you choose to leave, your activities will not be counted towards a Local Legend achievement.\n\nIf you'd like to join again in the future, only activities from the time you join will be counted.",
                "action_confirmation": {
                    "title": "Leave Local Legends",
                    "body": "If you choose to leave, your activites will not be counted towards a Local Legend achievement. It may take up to 24 hours for your achievements to be removed. If you'd like to join again in the future, only activities from the day you join will be counted. Are you sure you want to leave?",
                    "cancel": "Cancel",
                    "confirm": "Yes, Leave",
                    "action": "opt_out"
                }
            }
        }
    }
]
```

It looks like it will be difficult to determine which segments the user is KOM of. There is possibly this approach:
https://groups.google.com/g/strava-api/c/FNJSLI_HvLA/m/IsCbOxbyBwAJ
`curl https://www.strava.com/api/v3/athletes/2021127/koms`
however, it may be that it only shows LL of segments a person has created. If that's true, we will need to fallback to
looking through all segments they have starred, and then checking the LL.

`curl https://www.strava.com/api/v3/segments/starred`

In the response you get this, notice the is_kom field:
```
"athlete_pr_effort": {
            "id": 2941548630396809328,
            "activity_id": 6867780040,
            "elapsed_time": 145,
            "distance": 1485.64,
            "start_date": "2022-03-23T00:23:55Z",
            "start_date_local": "2022-03-22T17:23:55Z",
            "is_kom": false
        },
```


## inspiration projects
http://www.raceleap.live/
https://stravanity.vercel.app/
