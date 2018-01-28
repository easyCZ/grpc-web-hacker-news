import * as React from 'react';
import { Item } from 'semantic-ui-react';
import {Item as Story } from './proto/hackernews_pb';

type StoryListProps = {
  stories: Story.AsObject[],
  onStorySelect: (id: number) => void
};

const StoryList: React.SFC<StoryListProps> = (props) => {
  return (
    <Item.Group>
      {props.stories.map((story, i) =>
        <Item
          key={i}
          onClick={() => {
            if (story.id && story.id.id) {
              props.onStorySelect(story.id.id);
            }
          }}
        >
          <Item.Content>
            <Item.Header as="a">{story.title}</Item.Header>
            <Item.Meta>By: {story.by}</Item.Meta>
            <Item.Extra>{story.score} {story.time}</Item.Extra>
          </Item.Content>
        </Item>
      )}
    </Item.Group>
  );
};

export default StoryList;
