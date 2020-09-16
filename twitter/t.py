#!/usr/bin/env python3

import os
import requests
from pprint import pprint

tweet_url = 'https://twitter.com/kyxvbx364872/status/1228935216544112640?s=20'

parts = tweet_url.split('/')
username = parts[3]
tweet_id = parts[-1].split('?')[0]
proxies = { "http": "http://127.0.0.1:1081", "https": "http://127.0.0.1:1081", } 


headers = {
    'authority': 'api.twitter.com',
    'pragma': 'no-cache',
    'cache-control': 'no-cache',
    'x-twitter-client-language': 'zh-cn',
    'x-csrf-token': 'cea73f3255f2552e9cac584a151c9f59',
    'authorization': 'Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA',
    'user-agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.106 Safari/537.36',
    'sec-fetch-dest': 'empty',
    'x-twitter-auth-type': 'OAuth2Session',
    'x-twitter-active-user': 'yes',
    'accept': '*/*',
    'origin': 'https://twitter.com',
    'sec-fetch-site': 'same-site',
    'sec-fetch-mode': 'cors',
    'referer': 'https://twitter.com/pinngeeer678/status/1228923122063011842?s=20',
    'accept-language': 'zh-CN,zh;q=0.9,en;q=0.8',
    'cookie': 'personalization_id="v1_TUXIKoXDaiR5wUHaKOZFEQ=="; syndication_guest_id=v1%3A155306638976626710; guest_id=v1%3A155306639545558974; _ga=GA1.2.1298122637.1554346334; dnt=1; ads_prefs="HBESAAA="; kdt=JhgMsDc0xJgzbZ5wrJjMUHmlWE61b8yKVXS4S55u; remember_checked_on=1; csrf_same_site_set=1; csrf_same_site=1; auth_token=f29f1a5ba23da2407004470ffa4e4f6d184dfb2e; rweb_optin=side_no_out; twid=u%3D964887125773598720; tfw_exp=0; des_opt_in=Y; _gid=GA1.2.1645293137.1581834260; ct0=cea73f3255f2552e9cac584a151c9f59; _gat=1; external_referer=8e8t2xd8A2w%3D||4abf247TNXg4Rylyqt4v49u2LWyy1%2FaFyfd602NkflM%3D',
}

params = (
    ('include_profile_interstitial_type', '1'),
    ('include_blocking', '1'),
    ('include_blocked_by', '1'),
    ('include_followed_by', '1'),
    ('include_want_retweets', '1'),
    ('include_mute_edge', '1'),
    ('include_can_dm', '1'),
    ('include_can_media_tag', '1'),
    ('skip_status', '1'),
    ('cards_platform', 'Web-12'),
    ('include_cards', '1'),
    ('include_composer_source', 'true'),
    ('include_ext_alt_text', 'true'),
    ('include_reply_count', '1'),
    ('tweet_mode', 'extended'),
    ('include_entities', 'true'),
    ('include_user_entities', 'true'),
    ('include_ext_media_color', 'true'),
    ('include_ext_media_availability', 'true'),
    ('send_error_codes', 'true'),
    ('simple_quoted_tweets', 'true'),
    ('count', '20'),
    ('ext', 'mediaStats,cameraMoment'),
)

response = requests.get(f'https://api.twitter.com/2/timeline/conversation/{tweet_id}.json', headers=headers, params=params, proxies=proxies)
resp_json = response.json()

tweet                   = resp_json['globalObjects']['tweets'][str(tweet_id)]
user_id_str             = tweet['user_id_str']
full_text               = tweet['full_text']
user                    = resp_json['globalObjects']['users'][user_id_str]
name                    = user['name']
screen_name             = user['screen_name']
profile_image_url_https = user['profile_image_url_https']
profile_banner_url      = user['profile_banner_url']
medias                  = tweet['entities']['media']

print(f"user_id_str={user_id_str}")
print(f"screen_name={username}")
print(f"username={username}")
print(f"full_text={full_text}")
print(f"tweet_id={tweet_id}")


def local_path(url):
    parts = url.split('/')
    file_dir = '/'.join(parts[3:-1])
    file_name = parts[-1]
    file_path = os.path.join(file_dir, file_name)
    return file_dir, file_name, file_path


def download_media(url):
    print("download media: ", url)
    dir, name, path = local_path(url)
    if not os.path.exists(dir):
        os.makedirs(dir)

    if os.path.exists(path):
        return
        
    response = requests.get(url, proxies=proxies)
    if response.status_code != 200:
        return

    with open(os.path.join(dir, name), 'wb') as f:
        f.write(response.content)

    
download_media(profile_banner_url)
download_media(profile_image_url_https)
for media in medias:
    download_media(media['media_url_https'])
